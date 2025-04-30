package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

func GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is posts endpoint",
	})
}
