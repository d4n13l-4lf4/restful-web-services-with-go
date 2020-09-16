package main

import (
	"encoding/json"
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter09/longRunningTaskV1/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter09/longRunningTaskV1/models"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type JobServer struct {
	Queue amqp.Queue
	Channel *amqp.Channel
	Conn *amqp.Connection
}

func (s *JobServer) publish(jsonBody []byte) error {
	message := amqp.Publishing{
		ContentType: "application/json",
		Body: jsonBody,
	}
	err := s.Channel.Publish(
		"",
		queueName,
		false,
		false,
		message,
	)

	utils.HandleError(err, "Error while generating JobID")
	return err
}

func (s *JobServer) asyncDBHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()
	queryParams := r.URL.Query()
	unixTime, err := strconv.ParseInt(queryParams.Get("client_time"), 10, 64)
	clientTime := time.Unix(unixTime, 0)
	utils.HandleError(err, "Error while converting client time")

	jsonBody, err := json.Marshal(models.Job{
		ID: jobID,
		Type: "A",
		ExtraData: models.Log{ClientTime: clientTime},
	})

	utils.HandleError(err, "JSON body creation failed")

	if s.publish(jsonBody) == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}