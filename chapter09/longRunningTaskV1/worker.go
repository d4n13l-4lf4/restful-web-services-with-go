package main

import (
	"encoding/json"
	"time"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter09/longRunningTaskV1/models"
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter09/longRunningTaskV1/utils"
	"github.com/emicklei/go-restful/log"
	"github.com/streadway/amqp"
)

type Workers struct {
	conn *amqp.Connection
}

func (w *Workers) dbWork(job models.Job) {
	result := job.ExtraData.(map[string] interface{})
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