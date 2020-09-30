package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取文章列表
func GetList(c *gin.Context) {
	c.String(http.StatusOK, "GetList")
}

//新建文章
func Add(c *gin.Context) {
	c.String(http.StatusOK, "Add")
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
