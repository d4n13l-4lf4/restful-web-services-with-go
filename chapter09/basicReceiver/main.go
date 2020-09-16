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

func readQueue(done <-chan interface{}, message <-chan amqp.Delivery) {
	for {
		select {
		case <- done:
			log.Println("Exiting go routine")
			return
		case received := <- message:
			log.Printf("Received a message from the queue: %s", received.Body)
		}
	}
}

func main() {
	err := godotenv.Load()
	handleError(err, "Could not get env vars")
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
		false, /// Exclusive
		false, // No Waiting time
		nil, // Extra args
	)
	handleError(err, "Queue creation failed")
	messages, err := channel.Consume(
		testQueue.Name, // queue,
		"", // consumer
		true, // auto-acknowledge
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	handleError(err, "Failed to register a consumer")

	done := make(chan interface{})
	wait := make(chan interface{})
	go readQueue(done, messages)

	time.AfterFunc(10 * time.Second, func() {
		close(done)
		close(wait)
	})

	<- wait
}