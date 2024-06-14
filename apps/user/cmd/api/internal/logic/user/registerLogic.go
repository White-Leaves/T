package user

import (
	"context"
	"gomall/apps/user/cmd/api/internal/svc"
	"gomall/apps/user/cmd/api/internal/types"
	"gomall/apps/user/cmd/rpc/user"
	"gomall/apps/user/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {

	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		Nickname: req.Nickname,
		AuthKey:  req.Mobile,
		AuthType: model.UserAuthTypeSystem,
	})

	if err != nil {
		return nil, err
	}

	var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)

	return &resp, nil
}
