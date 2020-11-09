package middleware

import (
	"backend/utils"
	response "backend/utils/app"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetailed(response.AuthFailed, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(response.AuthExpired, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(response.AuthFailed, err.Error(), c)
			c.Abort()
			return
		}

		if claims.ExpiresAt-time.Now().Unix() < 60*60*24 {
			claims.ExpiresAt = time.Now().Unix() + 60*60*24*7
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
