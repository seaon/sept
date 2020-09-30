package v1

import (
	"backend/model"
	"backend/service"
	response "backend/utils/app"
	"fmt"
	"github.com/gin-gonic/gin"
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

	response.OkWithData(model.UserData{
		ID:       ret.ID,
		Username: ret.Username,
		MobileNo: ret.MobileNo,
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

	response.OkWithData(model.UserData{
		ID:       ret.ID,
		Username: ret.Username,
		MobileNo: ret.MobileNo,
	}, c)
	return
}
