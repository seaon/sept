package main

import (
	"fmt"

	"backend/config"
	"backend/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(config.GetServer().AppMode)

	port := fmt.Sprintf(":%d", config.GetServer().Port)

	route := router.InitRouter()

	route.Run(port) // listen and serve on 0.0.0.0:8080
}
