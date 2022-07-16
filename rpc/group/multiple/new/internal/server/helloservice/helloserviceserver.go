// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

package server

import (
	"context"

	"new/internal/logic/helloservice"
	"new/internal/svc"
	"new/pb"
)

type HelloServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedHelloServiceServer
}

func NewHelloServiceServer(svcCtx *svc.ServiceContext) *HelloServiceServer {
	return &HelloServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *HelloServiceServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	l := helloservicelogic.NewSayHelloLogic(ctx, s.svcCtx)
	return l.SayHello(in)
}