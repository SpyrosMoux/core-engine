package main

import (
	"github.com/spyrosmoux/core-engine/internal/logger"
	"github.com/spyrosmoux/core-engine/internal/pipelines"
	"github.com/spyrosmoux/core-engine/internal/queue"
)

func main() {
	// Setup custom Logger
	logger.Init()

	// Init RabbitMQ
	msgs := queue.InitRabbitMQ()

	var forever chan struct{}

	go func() {
		for d := range msgs {
			logger.Log(logger.InfoLevel, "Received a message with correlation id: "+d.CorrelationId)
			pipelines.RunJob(string(d.Body))
		}
	}()

	logger.Log(logger.InfoLevel, " [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
