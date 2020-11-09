package v1

import (
	"backend/model"
	"backend/service"
	"backend/utils"
	response "backend/utils/app"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取文章列表
func GetList(c *gin.Context) {
	c.String(http.StatusOK, "GetList")
}

//新建文章
func Add(c *gin.Context) {
	var article model.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		response.FailWithCode(response.InvalidArgument, fmt.Sprintf("%v", err), c)
		return
	}

	claims, _ := c.Get("claims")
	waitUse := claims.(*utils.CustomClaims)

	article.Uid = waitUse.ID

	err, id := service.Add(article)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), c)
		return
	}
	response.OkWithData(id, c)
}

//删除指定文章
func Del(c *gin.Context) {
	c.String(http.StatusOK, "Del")
}

//更新指定文章
func Update(c *gin.Context) {
	c.String(http.StatusOK, "Update")
}

//获取指定文章
func Get(c *gin.Context) {
	c.String(http.StatusOK, "Get")
}
