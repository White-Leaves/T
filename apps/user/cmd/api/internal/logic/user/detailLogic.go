package user

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"gomall/apps/user/cmd/api/internal/svc"
	"gomall/apps/user/cmd/api/internal/types"
	"gomall/apps/user/cmd/rpc/user"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {

	//userId := ctxdata.GetUidFromCtx(l.ctx)
	userInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	var userInfo types.User
	_ = copier.Copy(&userInfo, userInfoResp.User)

	return &types.UserInfoResp{
		UserInfo: userInfo,
	}, nil
}
