package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/d4n13l-4lf4/restful-web-services-with-go/chapter06/serverPush/protofiles"
)

const (
	address = "localhost:50051"
)

func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}

	c := pb.NewMoneyTransactionClient(conn)
	txRequest := &pb.TransactionRequest{
		To: "Me",
		From: "Me",
		Amount: 1,
	}
	ReceiveStream(c, txRequest)
}