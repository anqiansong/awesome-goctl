// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

package helloservice

import (
	"context"

	"hello/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	HelloReply   = pb.HelloReply
	HelloRequest = pb.HelloRequest

	HelloService interface {
		SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	}

	defaultHelloService struct {
		cli zrpc.Client
	}
)

func NewHelloService(cli zrpc.Client) HelloService {
	return &defaultHelloService{
		cli: cli,
	}
}

func (m *defaultHelloService) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	client := pb.NewHelloServiceClient(m.cli.Conn())
	return client.SayHello(ctx, in, opts...)
}
