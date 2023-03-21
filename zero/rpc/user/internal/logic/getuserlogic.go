package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-rest/zero/rpc/user/internal/svc"
	"go-rest/zero/rpc/user/types/pb"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.UserRequest) (*pb.UserResponse, error) {

	res, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:   res.UserId,
		Name: res.Username,
	}, nil
}
