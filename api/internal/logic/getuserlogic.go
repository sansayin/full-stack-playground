package logic

import (
	"context"

	"go-rest/api/internal/svc"
	"go-rest/api/internal/types"
	"go-rest/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
  res,err:=l.svcCtx.User.GetUser(l.ctx,&pb.UserRequest{
    Id:req.Id,
  })
  if(err!=nil){
    return nil, err
  }
  return &types.UserResponse{Id: req.Id, Name: res.Data.Name}, nil

}
