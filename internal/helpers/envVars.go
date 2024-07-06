package helpers

import (
	"github.com/joho/godotenv"
	"os"
	"spyrosmoux/core-engine/internal/logger"
)

func LoadEnvVariable(variable string) string {
	if err := godotenv.Load(); err != nil {
		logger.Log(logger.InfoLevel, "No .env file found, attempting to read from host environment variables")
	}

	variableValue := getEnvOrExit(variable)

	return variableValue
}

func getEnvOrExit(key string) string {
	value := os.Getenv(key)
	if value == "" {
		logger.Log(logger.FatalLevel, "Environment variable %s not set"+key)
	}
	return value
}
