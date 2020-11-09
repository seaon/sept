package v1

import (
	"backend/config"
	"backend/model"
	"backend/service"
	resp "backend/structures/response"
	"backend/utils"
	response "backend/utils/app"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func Login(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithCode(response.InvalidArgument, fmt.Sprintf("%v", err), c)
		return
	}

	err, ret := service.Login(user)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
		return
	}

	j := utils.NewJWT()
	token, _ := j.CreateToken(utils.CustomClaims{
		ID:       ret.ID,
		Username: ret.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 7天
			Issuer:    config.GetApp().JwtSecret,      // 签名的发行者
		},
	})

	response.OkWithData(resp.Login{
		Token: token,
		User: resp.UserData{
			ID:       ret.ID,
			Username: ret.Username,
		},
	}, c)
	return
}

func Logout(c *gin.Context) {
	response.Ok(c)
}

func Register(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithCode(response.InvalidArgument, fmt.Sprintf("%v", err), c)
		return
	}

	err, ret := service.Register(user)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
		return
	}

	response.OkWithData(resp.UserData{
		ID:       ret.ID,
		Username: ret.Username,
	}, c)
	return
}
