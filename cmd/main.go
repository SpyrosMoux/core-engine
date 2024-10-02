package main

import (
	"github.com/spyrosmoux/api/pkg/pipelineruns"
	"github.com/spyrosmoux/core-engine/internal/common"
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

	client := pipelineruns.NewClient(common.ApiBaseUrl)

	go func() {
		for d := range msgs {
			logger.Log(logger.InfoLevel, "Received a message with correlation id: "+d.CorrelationId)

			_, err := client.UpdatePipelineRunStatus(d.CorrelationId, "Running")
			if err != nil {
				logger.Log(logger.ErrorLevel, "Failed to update pipeline with error: "+err.Error())
			}

			runResult := true
			err = pipelines.RunPipeline(string(d.Body))
			if err != nil {
				runResult = false
				logger.Log(logger.ErrorLevel, "Failed to run pipeline with error: "+err.Error())
			}

			// Acknowledge the message after successful processing
			err = d.Ack(false)
			if err != nil {
				logger.Log(logger.ErrorLevel, "Failed to acknowledge message: "+err.Error())
			}

			if runResult {
				_, err = client.UpdatePipelineRunStatus(d.CorrelationId, "Completed")
				if err != nil {
					logger.Log(logger.ErrorLevel, "Failed to update pipeline with error: "+err.Error())
				}
			}

			_, err = client.UpdatePipelineRunStatus(d.CorrelationId, "Failed")
			if err != nil {
				logger.Log(logger.ErrorLevel, "Failed to update pipeline with error: "+err.Error())
			}
		}
	}()

	logger.Log(logger.InfoLevel, " [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
