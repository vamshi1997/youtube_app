package router

import (
	"github/vamshi1997/projects/go/youtube_app/internal/controllers"

	"github.com/gin-gonic/gin"
)

func InitAppRoutes(router *gin.Engine) {
	router.GET("/status", controllers.Status)
	router.GET("/fetch", controllers.FetchVideos)
	router.GET("/search", controllers.SearchVideos)
}
