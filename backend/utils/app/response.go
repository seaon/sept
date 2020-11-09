package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS         = 0
	ERROR           = 10
	AuthFailed      = 20
	AuthExpired     = 21
	InvalidArgument = 100
)

var empty map[string]string

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, empty, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, empty, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, empty, "失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, empty, message, c)
}

func FailWithCode(code int, message string, c *gin.Context) {
	Result(code, empty, message, c)
}

func FailWithDetailed(code int, message string, c *gin.Context) {
	Result(code, empty, message, c)
}
