package routes

import (
	"go-cms/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	r.GET("/", controllers.GetHome)
	api.GET("/posts", controllers.GetPosts)
}
