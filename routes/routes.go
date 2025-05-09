package routes

import (
	"go-cms/controllers"
	"go-cms/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	registerHomeRoutes(r)
	registerAPIRoutes(r, db)
	registerAuthRoutes(r, db)
}

func registerHomeRoutes(r *gin.Engine) {
	r.GET("/", controllers.GetHome)
}

func registerAPIRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1")

	postRoutes := api.Group("/posts")
	postRoutes.GET("/all", controllers.GetPosts)

	userRoutes := api.Group("/users")
	userRoutes.Use(middlewares.JWTAuthMiddleware())
	userRoutes.GET("/:user", func(c *gin.Context) {
		controllers.GetUser(c, db)
	})
}

func registerAuthRoutes(r *gin.Engine, db *gorm.DB) {
	auth := r.Group("/auth")
	auth.POST("/login", func(c *gin.Context) {
		controllers.Login(c, db)
	})
	auth.POST("/register", func(c *gin.Context) {
		controllers.RegisterUser(c, db)
	})
	auth.POST("/verify", func(c *gin.Context) {
		controllers.VerifyUser(c, db)
	})
}
