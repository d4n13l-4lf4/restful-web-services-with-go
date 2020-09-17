package main

import (
	"context"
	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter11/asyncClient/proto"
	"github.com/micro/go-micro"
	"log"
)

func ProcessEvent(ctx context.Context, event *proto.Event) error {
	log.Println("Got alert:", event)
	return nil
}

func main() {
	service := micro.NewService(micro.Name("weather_client"))
	service.Init()
	micro.RegisterSubscriber("alerts", service.Server(), ProcessEvent)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}