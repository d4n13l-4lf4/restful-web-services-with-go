package main

import (
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main(){
	err := godotenv.Load()
	handleError(err, "Couldn't read env vars")

	connectionString := os.Getenv("AMQP_CONNECTION_STRING")
	conn, err := amqp.Dial(connectionString)
	handleError(err, "Dialing failed to RabbitMQ broker")
	defer conn.Close()

	channel, err := conn.Channel()
	handleError(err, "Fetching channel failed")
	defer channel.Close()

	testQueue, err := channel.QueueDeclare(
		"test", // Name of the queue
		false, // Message is persisted or not
		false, // Delete message when unused
		false, // Exclusive
		false, // No Waiting time
		nil, // Extra args
	)

	handleError(err, "Queue creation failed")

	serverTime := time.Now()
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte(serverTime.String()),
	}
	err = channel.Publish(
		"", // exchange
		testQueue.Name, // routing key(Queue)
		false, // mandatory
		false, // immediate
		message,
	)
	handleError(err, "Failed to publish a message")
	log.Println("Successfully published a message to the queue")
}