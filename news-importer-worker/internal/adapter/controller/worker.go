package controller

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"net/url"
	"silver.com/internal/entity"
	"silver.com/internal/infra/broker"
	"time"
)

type WorkerController struct {
	broker *broker.RabbitMQ
	ok     int64
	fail   int64
}

func NewWorkerController(broker *broker.RabbitMQ) *WorkerController {
	return &WorkerController{
		broker: broker,
		ok:     0,
		fail:   0,
	}
}

func (s *WorkerController) Run() {
	messages := s.setConsumer()
	log.Println("waiting for messages")

	go s.metrics()

	for message := range messages {
		go s.process(message)
		message.Ack(true)
	}
}

func (s *WorkerController) process(message amqp091.Delivery) {
	job, _ := entity.Deserialize(message.Body)

	if _, err := url.ParseRequestURI(job.Arguments); err != nil {
		log.Printf("WARNING. Job name: %s, url: %s. Invalid URI", job.Name, job.Arguments)
		return
	}

	_, err := http.Get(job.Arguments)
	if err != nil {
		s.fail += 1
		log.Printf("ERROR. Job name: %s, url: %s. %s", job.Name, job.Arguments, err.Error())
	} else {
		s.ok += 1
	}
}

func (s *WorkerController) setConsumer() <-chan amqp091.Delivery {
	err := s.broker.Channel.Qos(4, 0, false)
	if err != nil {
		log.Fatal(err)
	}

	messages, err := s.broker.Channel.Consume(broker.QUEUE, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	return messages
}

func (s *WorkerController) metrics() {
	for {
		log.Printf("ok: %d, fail: %d", s.ok, s.fail)
		time.Sleep(30 * time.Second)
	}
}
