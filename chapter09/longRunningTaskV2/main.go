package main

import (
	"github.com/go-redis/redis"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"github.com/joho/godotenv"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter09/longRunningTaskV2/utils"
)

const (
	queueName = "jobQueue"
	hostString = "127.0.0.1:8000"
)

func getServer(name string) JobServer {
	connectionString := os.Getenv("AMQP_CONNECTION_STRING")
	conn, err := amqp.Dial(connectionString)
	utils.HandleError(err, "Dialing failed to RabbitMQ broker")

	channel, err := conn.Channel()
	utils.HandleError(err, "Fetching channel failed")

	jobQueue, err := channel.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	utils.HandleError(err, "Job queue creation failed")
	return JobServer{Conn: conn, Channel: channel, Queue: jobQueue}
}

func main() {
	err := godotenv.Load()
	utils.HandleError(err, "Could not load env vars")

	jobServer := getServer(queueName)
	jobServer.redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	go func(conn *amqp.Connection) {
		workerProcess := Workers{
			conn: jobServer.Conn,
		}
		workerProcess.run()
	} (jobServer.Conn)

	router := mux.NewRouter()
	router.HandleFunc("/job/database", jobServer.asyncDBHandler)
	router.HandleFunc("/job/status", jobServer.statusHandler)
	// router.HandleFunc("/job/mail", jobServer.asyncMailHandler)
	// router.HandleFunc("job/callback", jobServer.asyncCallbackHandler)

	srv := http.Server{
		Handler: router,
		Addr: hostString,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	defer jobServer.Channel.Close()
	defer jobServer.Conn.Close()
}