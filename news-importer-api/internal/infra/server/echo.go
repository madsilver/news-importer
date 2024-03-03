package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"silver.com/internal/infra/env"
)

type EchoServer struct {
}

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

func (s *EchoServer) Start(manager *Manager) {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/v1/jobs", manager.JobController.FindJobs)

	e.Logger.Fatal(e.Start(":" + env.GetString("SERVER_PORT", "8000")))
}
