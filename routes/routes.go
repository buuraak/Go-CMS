package routes

import (
	"go-cms/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	registerHomeRoutes(r)
	registerAPIRoutes(r, db)
}

func registerHomeRoutes(r *gin.Engine) {
	r.GET("/", controllers.GetHome)
}

func registerAPIRoutes(r *gin.Engine, db *gorm.DB) {

	api := r.Group("/api/v1")
	// post endpoints
	api.GET("/posts", controllers.GetPosts)

	// User endpoints
	api.GET("/users/:user", func(c *gin.Context) {
		controllers.GetUser(c, db)
	})
}
