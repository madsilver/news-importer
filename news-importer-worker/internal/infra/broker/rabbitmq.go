package broker

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"silver.com/internal/infra/env"
)

const QUEUE = "schedules"

type RabbitMQ struct {
	Channel *amqp091.Channel
}

func NewRabbitMQ() *RabbitMQ {
	conn, err := amqp091.Dial(getDSN())
	if err != nil {
		log.Fatal(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(QUEUE, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connected to RabbitMQ")

	return &RabbitMQ{
		Channel: channel,
	}
}

func getDSN() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s",
		env.GetString("RABBITMQ_USER", "sapo"),
		env.GetString("RABBITMQ_PASSWORD", "sapo"),
		env.GetString("RABBITMQ_HOST", "localhost"),
		env.GetString("RABBITMQ_PORT", "5672"))
}
