package controllers

import (
	"go-cms/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func GetUser(c *gin.Context, db *gorm.DB) {
	uid := c.Param("user")
	user := models.User{}
	err := db.Where("id = ? OR username = ? OR email = ?", uid, uid, uid).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"Not found": "User not found"})
		} else {
			log.Fatalf("Error fetching user %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error while retrieving user"})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}
