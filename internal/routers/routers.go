package routers

import (
	"github.com/gin-gonic/gin"
	"spyrosmoux/core-engine/internal/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/webhook", handlers.HandleWebhook)

	return router
}
