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
	// todo: add your logic here and delete this line

	return &pb.ListResponse{}, nil
}
