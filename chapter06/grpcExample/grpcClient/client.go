package main

import (
	"log"

	pb "github.com/d4n13l-4lf4/restful-web-services-with-go/chapter06/grpcExample/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error: %v", err)
		panic(err)
	}

	c := pb.NewMoneyTransactionClient(conn)
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from, To: to, Amount: amount})
	log.Printf("Transaction confirmed: %v\n", r.Confirmation)
}