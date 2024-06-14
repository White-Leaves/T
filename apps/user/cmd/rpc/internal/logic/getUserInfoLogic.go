package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"gomall/apps/user/cmd/rpc/internal/svc"
	"gomall/apps/user/cmd/rpc/pb"
	"gomall/apps/user/cmd/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	tuser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if tuser == nil {
		return nil, err
	}

	var respUser user.User
	_ = copier.Copy(&respUser, tuser)

	return &user.GetUserInfoResp{
		User: &respUser,
	}, nil
}
