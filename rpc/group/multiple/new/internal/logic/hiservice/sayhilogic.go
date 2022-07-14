package hiservicelogic

import (
	"context"

	"new/internal/svc"
	"new/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHiLogic {
	return &SayHiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHiLogic) SayHi(in *pb.HelloRequest) (*pb.HelloReply, error) {
	// todo: add your logic here and delete this line

	return &pb.HelloReply{}, nil
}
