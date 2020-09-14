package main

import (
	"fmt"
	"encoding/json"

	pb "github.com/d4n13l-4lf4/restful-web-services-with-go/chapter06/protobufs/protofiles"
)

func main() {
	p := &pb.Person{
		Id: 1234,
		Name: "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	body, _ := json.Marshal(p)
	fmt.Println(string(body))
}