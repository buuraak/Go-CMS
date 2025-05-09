package controllers

import (
	"go-cms/config"
	"go-cms/helpers"
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

	if err := helpers.CheckPassword(user.Password, input.Password); err != nil {
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

func RegisterUser(c *gin.Context, db *gorm.DB) {
	var input struct {
		Username  string `json:"username" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Email     string `json:"email" binding:"required"`
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Role      string `json:"role" binding:"required,oneof=admin editor customer"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	var user models.User

	if err := db.Where("username = ?", input.Username).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error whilst checking for username"})
		return
	}

	if err := db.Where("email = ?", input.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error whilst checking for email"})
		return
	}

	verificationToken, err := helpers.GenerateVerificationToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate verification token"})
	}

	user = models.User{
		Username:          input.Username,
		Password:          hashedPassword,
		Email:             input.Email,
		FirstName:         input.FirstName,
		LastName:          input.LastName,
		Role:              input.Role,
		VerificationToken: verificationToken,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered succesfully"})
}

func VerifyUser(c *gin.Context, db *gorm.DB) {
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "verification token is missing"})
		return
	}

	var user models.User
	if err := db.Where("verification_token = ?", token).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error whilst trying to verify user"})
			return
		}
	}

	user.IsVerified = true
	user.VerificationToken = ""
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User verified successfully"})
}
