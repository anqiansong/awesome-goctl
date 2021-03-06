// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

package server

import (
	"context"

	"client/hello"
	"client/internal/logic"
	"client/internal/svc"
)

type HelloServer struct {
	svcCtx *svc.ServiceContext
	hello.UnimplementedHelloServer
}

func NewHelloServer(svcCtx *svc.ServiceContext) *HelloServer {
	return &HelloServer{
		svcCtx: svcCtx,
	}
}

func (s *HelloServer) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	l := logic.NewSayHelloLogic(ctx, s.svcCtx)
	return l.SayHello(in)
}
