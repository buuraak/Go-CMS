package controllers

import (
	"go-cms/config"
	"go-cms/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	claims, _ := c.Get("claims")
	userClaims := claims.(jwt.MapClaims)

	if userClaims["role"] != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Acces denied"})
		return
	}

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

func Login(c *gin.Context, db *gorm.DB) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	var user models.User
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	secret := config.GetJWTSecret()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
