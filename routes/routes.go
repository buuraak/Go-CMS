package routes

import (
	"go-cms/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	registerHomeRoutes(r)
	registerAPIRoutes(r)
}

func registerHomeRoutes(r *gin.Engine) {
	r.GET("/", controllers.GetHome)
}

func registerAPIRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.GET("/posts", controllers.GetPosts)
}
