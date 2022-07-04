package logic

import (
	"context"
	"fmt"

	"tag/internal/svc"
	"tag/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHeaderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHeaderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHeaderLogic {
	return &GetHeaderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHeaderLogic) GetHeader(header types.UserHeader) error {
	// todo: add your logic here and delete this line
	fmt.Println("header:", header.Authorization)
	return nil
}
