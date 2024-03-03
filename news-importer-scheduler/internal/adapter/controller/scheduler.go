package controller

import (
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"silver.com/internal/entity"
	"silver.com/internal/infra/broker"
	"silver.com/internal/infra/env"
)

const INTERVAL = "@every 30s"

type SchedulerController struct {
	broker *broker.RabbitMQ
}

func NewSchedulerController(broker *broker.RabbitMQ) *SchedulerController {
	return &SchedulerController{
		broker: broker,
	}
}

func (s *SchedulerController) Run() {
	c := cron.New()
	defer c.Stop()

	err := c.AddFunc(INTERVAL, s.GetJobs)

	if err != nil {
		log.Fatal(err.Error())
	}

	c.Start()
	log.Println("news-importer-scheduler running...")

	forever := make(chan bool)
	<-forever
	//select {}
}

func (s *SchedulerController) GetJobs() {
	jobs := fetchJobsFromServer()

	log.Printf("jobs found: %d", len(jobs))
	log.Println("sending messages to broker")

	for _, job := range jobs {
		s.SendMessage(job)
	}
}

func (s *SchedulerController) SendMessage(job *entity.Job) {
	encode, _ := job.Serialize()
	message := amqp091.Publishing{ContentType: "text/plain", Body: encode}

	err := s.broker.Channel.Publish("", broker.QUEUE, false, false, message)
	if err != nil {
		log.Print(err)
	}
}

func fetchJobsFromServer() []*entity.Job {
	url := getURI("jobs")
	log.Println("fetching jobs from " + url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	var jobs []*entity.Job
	if err := json.NewDecoder(resp.Body).Decode(&jobs); err != nil {
		log.Fatal(err.Error())
	}

	return jobs
}

func getURI(resource string) string {
	return fmt.Sprintf("http://%s:%s/v1/%s",
		env.GetString("API_HOST", "localhost"),
		env.GetString("API_PORT", "8000"),
		resource)
}
