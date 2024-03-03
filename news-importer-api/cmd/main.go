package main

import (
	"log"
	"silver.com/internal/adapter/controller"
	"silver.com/internal/adapter/repository/mysql"
	"silver.com/internal/infra/db"
	"silver.com/internal/infra/server"
)

func main() {
	log.Println("news-importer-api start")

	repository := mysql.NewJobRepository(db.NewMysqlDB())

	manager := &server.Manager{
		JobController: controller.NewJobController(repository),
	}

	server.NewEchoServer().Start(manager)
}
