package logic

import (
	"context"
	"go-rest/api/internal/svc"
	"go-rest/api/internal/types"
	"go-rest/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers() (resp []types.UserInfo, err error) {
	res, err := l.svcCtx.User.GetAllUsers(l.ctx, &pb.Empty{})
  if err!=nil {
    return nil, err
  }
	for _, user := range res.GetData() {
		resp = append(resp, types.UserInfo{
			Id:   user.Id,
			Name: user.Name,
		})
	}
	return resp, nil
}
