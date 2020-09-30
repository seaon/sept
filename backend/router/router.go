package router

import (
	v1 "backend/controller/v1"
	"backend/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	router.POST("/login", v1.Login)
	router.POST("/logout", v1.Logout)
	router.POST("/register", v1.Register)

	router.GET("/articles", v1.GetList)
	router.GET("/article/:id", v1.Get)

	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized := router.Group("")
	authorized.Use(middleware.Check())
	{
		authorized.POST("/article", v1.Add)
		authorized.PUT("/article/:id", v1.Update)
		authorized.DELETE("/article/:id", v1.Del)
	}

	return router
}
