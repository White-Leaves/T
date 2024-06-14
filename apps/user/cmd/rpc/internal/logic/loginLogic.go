package logic

import (
	"context"
	"github.com/pkg/errors"
	"gomall/apps/user/cmd/rpc/internal/svc"
	"gomall/apps/user/cmd/rpc/pb"
	"gomall/apps/user/cmd/rpc/user"
	"gomall/apps/user/model"
	"gomall/pkg/Md5"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {

	var userId int64
	var err error

	switch in.AuthType {
	case model.UserAuthTypeSystem:
		userId, err = l.loginByMobile(in.AuthKey, in.Password)
	default:
		return nil, errors.New("server error")
	}
	if err != nil {
		return nil, err
	}

	generateLoginToken := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateLoginToken.GenerateToken(&user.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &user.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, mobile)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.New("手机号查询用户失败")
	}
	if user == nil {
		return 0, errors.New("用户不存在")
	}

	if !(Md5.Md5encrypt(password) == user.Password) {
		return 0, errors.New("密码出错")
	}

	return user.Id, nil
}
