package middleware

import (
	"github.com/gin-gonic/gin"
)

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {

		//if code != e.SUCCESS {
		//    c.JSON(http.StatusUnauthorized, gin.H{
		//        "code": code,
		//        "msg":  e.GetMsg(code),
		//        "data": data,
		//    })
		//
		//    c.Abort()
		//    return
		//}
		c.Next()
	}
}
