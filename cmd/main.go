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
			err := pipelines.RunPipeline(string(d.Body))
			if err != nil {
				logger.Log(logger.ErrorLevel, "Failed to run pipeline with error: "+err.Error())
			}

			// Acknowledge the message after successful processing
			err = d.Ack(false)
			if err != nil {
				logger.Log(logger.ErrorLevel, "Failed to acknowledge message: "+err.Error())
			}
		}
	}()

	logger.Log(logger.InfoLevel, " [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
