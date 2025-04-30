package main

import (
	"go-cms/config"
	"go-cms/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":" + config.GetEnv("PORT", "8080"))
}
