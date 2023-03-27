package logic

import (
	"context"

	"go-rest/api/internal/svc"
	"go-rest/api/internal/types"
	"go-rest/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserInfo) (resp *types.UserResponse, err error) {
	res, err := l.svcCtx.User.UpdateUser(l.ctx, &pb.UserInfo{
		Id:   req.Id,
		Name: req.Name,
	})
  if(err!=nil){
    return nil, err
  }
	return &types.UserResponse{
   Id: res.Data.Id, 
   Name: res.Data.Name,
  }, err
}
