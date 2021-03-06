package main

import (
	"context"

	"github.com/d4n13l-4lf4/restful-web-services-with-go/chapter11/encryptService/proto"
)

type Encrypter struct {
}

/* func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req *proto.Request, res *proto.Response) error {
		fmt.Printf("encryption request at time: %v", time.Now())
		return fn(ctx, req, res)
	}
}  */

func (g *Encrypter) Encrypt(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Result = EncryptString(req.Key, req.Message)
	return nil
}

func (g *Encrypter) Decrypt(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Result = DecryptString(req.Key, req.Message)
	return nil
}
