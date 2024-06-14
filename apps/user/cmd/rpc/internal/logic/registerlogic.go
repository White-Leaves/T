package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gomall/apps/user/cmd/rpc/user"
	"gomall/apps/user/model"
	"gomall/pkg/Md5"

	"github.com/zeromicro/go-zero/core/logx"
	"gomall/apps/user/cmd/rpc/internal/svc"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	tuser, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(err, "model fail")
	}
	if tuser != nil {
		return nil, errors.Wrapf(err, "Register user exists mobile")
	}

	var userId int64
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		tuser := new(model.User)
		tuser.Mobile = in.Mobile
		tuser.Nickname = in.Nickname
		tuser.Password = in.Password

		if len(tuser.Nickname) == 0 {
			return errors.New("nickname is null")
		}
		if len(in.Password) > 0 {
			tuser.Password = Md5.Md5encrypt(in.Password)
		}
		insertResult, err := l.svcCtx.UserModel.InsertR(ctx, session, tuser)
		if err != nil {
			return errors.Wrapf(err, "Register db user Insert err:%v,user:%+v", err, tuser)
		}
		lastId, err := insertResult.LastInsertId()
		if err != nil {
			return errors.Wrapf(err, "Register db user insertResult.LastInsertId err:%v,user:%+v", err, insertResult)
		}
		userId = lastId

		userAuth := new(model.UserAuth)
		userAuth.UserId = lastId
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		if _, err := l.svcCtx.UserAuthModel.InsertR(ctx, session, userAuth); err != nil {
			return errors.Wrapf(err, "Register db user_auth Insert err:%v,userAuth:%v", err, userAuth)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	//2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "GenerateToken userId : %d", userId)
	}

	return &user.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil

}
