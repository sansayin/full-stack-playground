package logic

import (
	"context"

	"go-rest/rpc/internal/svc"
	"go-rest/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
  res, err := l.svcCtx.Model.FindOne(l.ctx,in.Id)
  if(err!=nil){
    return nil, err
  }
	// todo: add your logic here and delete this line

	return &pb.UserResponse{
    Data: &pb.UserInfo{
      Id: res.UserId,
      Name: res.Username,
    },
  }, nil
}
