package main

import (
	"log"
	"silver.com/internal/adapter/controller"
	"silver.com/internal/infra/broker"
)

func main() {
	log.Println("news-importer-worker start")

	rabbitmq := broker.NewRabbitMQ()

	controller.NewWorkerController(rabbitmq).Run()
}
