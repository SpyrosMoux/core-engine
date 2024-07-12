package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"spyrosmoux/core-engine/internal/common"
	"spyrosmoux/core-engine/internal/logger"
)

func InitRabbitMQ() <-chan amqp.Delivery {
	conn, err := amqp.Dial("amqp://" + common.RabbitMQUser + ":" + common.RabbitMQPassword + "@" + common.RabbitMQHost + ":" + common.RabbitMQPort + "/")
	if err != nil {
		logger.Log(logger.FatalLevel, "Failed to connect to RabbitMQ "+err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Log(logger.FatalLevel, "Failed to open a channel: "+err.Error())
	}

	q, err := ch.QueueDeclare(
		"jobs",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Log(logger.FatalLevel, "Failed to declare a queue: "+err.Error())
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Log(logger.FatalLevel, "Failed to register a consumer: "+err.Error())
	}

	return msgs
}
