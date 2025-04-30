package routes

import (
	"go-cms/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		// Home Route
		r.GET("/", controllers.GetHome)

		// API routes for posts
		api.GET("/posts", controllers.GetPosts)
	}
}
