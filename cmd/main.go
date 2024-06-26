package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"spyrosmoux/core-engine/internal/helpers"
	"spyrosmoux/core-engine/internal/models"
	"spyrosmoux/core-engine/internal/queue"
)

func main() {
	msgs := queue.InitRabbitMQ()

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message with correlation id: %s", d.CorrelationId)
			runJob(d)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func runJob(job amqp.Delivery) {
	var ci models.UnifiedCI
	err := helpers.ReadYAMLFromString(string(job.Body), &ci)
	failOnError(err, "Failed to read yaml")

	for _, job := range ci.Jobs {
		err := helpers.ExecuteJob(job, ci.Variables)
		if err != nil {
			log.Fatalf("Error executing job: %v", err)
		}
	}
}
