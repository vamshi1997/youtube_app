package router

import (
	"fmt"
	"github/vamshi1997/projects/go/youtube_app/internal/boot"

	"github.com/gin-gonic/gin"
)

func InitiateRouter() {
	router := gin.New()

	cfg := boot.GetConfig()

	InitAppRoutes(router)

	serverAddr := fmt.Sprintf("%s:%v", cfg.AppConfig.Server.Host, cfg.AppConfig.Server.Port)
	router.Run(serverAddr)

	fmt.Printf("Server Started Successfully & listening to Port: %v", cfg.AppConfig.Server.Port)
}
