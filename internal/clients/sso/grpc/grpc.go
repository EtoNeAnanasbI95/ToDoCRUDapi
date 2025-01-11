package grpc

import (
	authv1 "github.com/EtoNeAnanasbI95/protos_auth/gen/go"
	"google.golang.org/grpc"
)

type Client struct {
	Api authv1.AuthClient
	Con *grpc.ClientConn
}

func New(addr string) (*Client, error) {
	cc, err := grpc.NewClient(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := authv1.NewAuthClient(cc)
	return &Client{Api: client, Con: cc}, nil
}
