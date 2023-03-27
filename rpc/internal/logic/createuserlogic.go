package logic

import (
	"context"
	"go-rest/model"
	"go-rest/rpc/internal/svc"
	"go-rest/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *pb.UserInfo) (*pb.UserResponse, error) {
	user := &model.ZeroUser{
		Username: in.GetName(),
	}
	ret, err := l.svcCtx.Model.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}
	data := &pb.UserInfo{
		Id:   ret.UserId,
		Name: ret.Username,
	}
	return &pb.UserResponse{Data: data}, nil
}
