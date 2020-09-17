package main

import (
	"fmt"

	"github.com/micro/go-micro"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter11/encryptService/proto"
)

func main() {
	service := micro.NewService(
			micro.Name("encrypter"),
		)
	service.Init()

	proto.RegisterEncrypterHandler(service.Server(), new(Encrypter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
