package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-awesome/user/internal/svc"
	"go-zero-awesome/user/internal/types"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.Request) (resp *types.Response, err error) {
	return &types.Response{
		Message: "Hello " + req.Name,
	}, nil
}
