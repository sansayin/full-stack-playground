package logic

import (
	"context"

	"go-rest/rpc/internal/svc"
	"go-rest/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUsersLogic {
	return &GetAllUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllUsersLogic) GetAllUsers(in *pb.Empty) (*pb.ListResponse, error) {
  resp, err:=l.svcCtx.Model.FindMany(l.ctx)
  if(err!=nil){
    return nil, err
  }
  var data = make([]*pb.UserInfo,0)
  for _,user:=range(*resp){
    data = append(data, &pb.UserInfo{
      Id: user.UserId,
      Name: user.Username,
    })
  }
 	return &pb.ListResponse{
    Data: data,
  }, nil
}
