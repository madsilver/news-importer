package main

import (
	"log"
	"silver.com/internal/adapter/controller"
	"silver.com/internal/infra/broker"
)

func main() {
	log.Println("news-importer-scheduler start")

	rabbitmq := broker.NewRabbitMQ()

	controller.NewSchedulerController(rabbitmq).Run()
}
