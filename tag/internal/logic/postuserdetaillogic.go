package logic

import (
	"context"
	"fmt"

	"tag/internal/svc"
	"tag/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostUserDetailLogic {
	return &PostUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostUserDetailLogic) PostUserDetail(req *types.UserDetailReq) error {
	// todo: add your logic here and delete this line
	fmt.Println("post ID:", req.ID)
	fmt.Println("post Name:", req.Name)
	return nil
}
