package logic

import (
	"context"
	"fmt"
	"strconv"

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
	id, err := strconv.Atoi(in.GetId())
	if err != nil {
		return nil, nil
	}
	user := &model.ZeroUser{
		UserId:   int64(id),
		Username: in.GetName(),
	}
	_, err1 := l.svcCtx.Model.Insert(l.ctx, user)
	if err1 != nil {
		return nil, nil
	}
	data := &pb.UserInfo{
		Id:   fmt.Sprint(id),
		Name: in.GetName(),
	}
	return &pb.UserResponse{Data: data}, nil
}
