package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spyrosmoux/core-engine/internal/common"
	"github.com/spyrosmoux/core-engine/internal/logger"
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
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Log(logger.FatalLevel, "Failed to declare a queue: "+err.Error())
	}

	// Set QoS (Quality of Service) to prefetch 1 message at a time
	err = ch.Qos(1, 0, false)
	if err != nil {
		logger.Log(logger.ErrorLevel, "Failed to set QoS: "+err.Error())
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
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
