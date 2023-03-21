package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-rest/zero/api/internal/svc"
	"go-rest/zero/api/internal/types"
	"go-rest/zero/rpc/user/users"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserId) (resp *types.UserInfo, err error) {
	res, err := l.svcCtx.Users.GetUser(l.ctx, &users.UserRequest{
		Id: req.Id,
	})
	return &types.UserInfo{
		res.Id,
		res.Name,
	}, err
}
