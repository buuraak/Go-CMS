package main

import (
	"go-cms/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})

	r.Run(":" + config.GetEnv("PORT", "8080"))
}
