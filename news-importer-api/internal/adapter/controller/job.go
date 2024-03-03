package controller

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"silver.com/internal/adapter/repository/mysql"
	"silver.com/internal/entity"
)

type JobController struct {
	repository *mysql.JobRepository
}

func NewJobController(repository *mysql.JobRepository) *JobController {
	return &JobController{
		repository: repository,
	}
}

func (c *JobController) FindJobs(ctx echo.Context) error {
	jobs, err := c.repository.FetchJobs()
	if err != nil {
		log.Println("Error: " + err.Error())
		return ctx.JSON(http.StatusInternalServerError, entity.Response{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, jobs)
}
