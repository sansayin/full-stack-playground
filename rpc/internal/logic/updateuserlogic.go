package logic

import (
	"context"

	"go-rest/model"
	"go-rest/rpc/internal/svc"
	"go-rest/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *pb.UserInfo) (*pb.UserResponse, error) {
	err := l.svcCtx.Model.Update(l.ctx, &model.ZeroUser{
		UserId:   in.Id,
		Username: in.Name,
	})
	if err != nil {
    return nil, err
	}

	return &pb.UserResponse{
    Data:&pb.UserInfo{
      Id: in.Id,
      Name: in.Name,
    },
  }, nil
}
