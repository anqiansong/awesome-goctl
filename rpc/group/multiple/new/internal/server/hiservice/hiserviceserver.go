// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

package server

import (
	"context"

	"new/internal/logic/hiservice"
	"new/internal/svc"
	"new/pb"
)

type HiServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedHiServiceServer
}

func NewHiServiceServer(svcCtx *svc.ServiceContext) *HiServiceServer {
	return &HiServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *HiServiceServer) SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	l := hiservicelogic.NewSayHiLogic(ctx, s.svcCtx)
	return l.SayHi(in)
}