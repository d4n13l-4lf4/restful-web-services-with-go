package main

import (
	"context"
	"fmt"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter11/encryptClient/proto"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("encrypter.client"))
	service.Init()

	encrypter := proto.NewEncrypterService("encrypter", service.Client())

	res, err := encrypter.Encrypt(context.TODO(), &proto.Request{
		Message: "I am a message",
		Key: "111023043350789514532147",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Result)

	res, err = encrypter.Decrypt(context.TODO(), &proto.Request{
		Message: res.Result,
		Key: "111023043350789514532147",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Result)
}
