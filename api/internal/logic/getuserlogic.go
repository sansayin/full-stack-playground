package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-rest/api/internal/svc"
	"go-rest/api/internal/types"
	"go-rest/rpc/pb"
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

func (l *GetUserLogic) GetUser(req *types.UserRequest) (resp *types.UserResponse, err error) {
	ret, err := l.svcCtx.User.GetUser(l.ctx, &pb.UserRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserResponse{Id: ret.Data.GetId(), Name: ret.Data.GetName()}, nil
}
