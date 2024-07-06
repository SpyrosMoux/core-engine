package main

import (
	"spyrosmoux/core-engine/internal/common"
	"spyrosmoux/core-engine/internal/logger"
	"spyrosmoux/core-engine/internal/routers"
)

func main() {
	// Setup custom Logger
	logger.Init()

	// Setup routes
	router := routers.SetupRouter()

	logger.Log(logger.InfoLevel, "Starting server on port "+common.ApiPort)

	err := router.Run(":" + common.ApiPort)
	if err != nil {
		logger.Log(logger.FatalLevel, err.Error())
	}
}
