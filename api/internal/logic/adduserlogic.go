package logic

import (
	"context"
	"log"

	"go-rest/api/internal/svc"
	"go-rest/api/internal/types"
	"go-rest/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.UserInfo) (resp *types.UserResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.User.CreateUser(l.ctx, &pb.UserInfo{
		Id:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		log.Printf("Error: %v\n", err)
		return nil, err
	}
	return &types.UserResponse{
		Id:   res.Data.Id,
		Name: res.Data.Name,
	}, err
}
