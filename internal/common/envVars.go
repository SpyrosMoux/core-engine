package common

import "github.com/spyrosmoux/core-engine/internal/helpers"

var (
	RabbitMQHost     = helpers.LoadEnvVariable("RABBITMQ_HOST")
	RabbitMQUser     = helpers.LoadEnvVariable("RABBITMQ_USER")
	RabbitMQPassword = helpers.LoadEnvVariable("RABBITMQ_PASSWORD")
	RabbitMQPort     = helpers.LoadEnvVariable("RABBITMQ_PORT")
	ApiBaseUrl       = helpers.LoadEnvVariable("API_BASE_URL")
)
