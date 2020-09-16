package main

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter09/longRunningTaskV2/models"
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter09/longRunningTaskV2/utils"
	"github.com/emicklei/go-restful/log"
	"github.com/streadway/amqp"
)

type Workers struct {
	conn *amqp.Connection
	redisClient *redis.Client
}

func (w *Workers) dbWork(job models.Job) {
	result := job.ExtraData.(map[string] interface{})
	w.redisClient.Set(job.ID.String(), "STARTED", 0)
	log.Printf("Worker %s: extracting data..., JOB: %s", job.Type, result)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s, saving data to database..., JOB: %s", job.Type, job.ID)
}

func (w *Workers) callbackWork(job models.Job) {
	log.Printf("Worker %s: performing some long running process..., JOB: %s", job.Type, job.ID)
	time.Sleep(10 * time.Second)
	log.Printf("Worker %s: posting the data back to he give callback..., JOB: %s", job.Type, job.ID)
}

func (w *Workers) emailWork(job models.Job) {
	log.Printf("Worker %s: sending the email..., JOB: %s", job.Type, job.ID)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s: sent the email successfully, JOB: %s", job.Type, job.ID)
}

func (w *Workers) run() {
	log.Printf("Workers are booted up and running")
	channel, err := w.conn.Channel()
	utils.HandleError(err, "Fetching channel failed")
	defer channel.Close()

	w.redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	jobQueue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.HandleError(err, "Job queue fetch failed")

	messages, err := channel.Consume(
		jobQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func () {
		for message := range messages {
			job := models.Job{}
			err = json.Unmarshal(message.Body, &job)
			log.Printf("Workers received a message from the queue: %s", job)
			utils.HandleError(err, "Unable to load queue message")

			switch job.Type {
			case "A":
				w.dbWork(job)
			case "B":
				w.callbackWork(job)
			case "C":
				w.emailWork(job)
			}
		}
	}()
	defer w.conn.Close()
	wait := make(chan bool)
	<-wait
}