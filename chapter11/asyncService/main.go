package main

import (
	"context"
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter11/asyncService/proto"
	"github.com/micro/go-micro"
	"log"
	"time"
)

func main() {
	service := micro.NewService(
		micro.Name("weather"),
	)
	p := micro.NewPublisher("alerts", service.Client())

	go func() {
		for now := range time.Tick(15 * time.Second) {
			log.Println("Publishing weather alert to Topic: alerts")
			p.Publish(context.Background(), &proto.Event{
				City:        "Quito",
				Timestamp:   now.UTC().Unix(),
				Temperature: 2,
			})
		}
	}()

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
