package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		InsertR(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *customUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultUserModel) InsertR(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	gomallUserIdKey := fmt.Sprintf("%s%v", cacheGomallUserUserMobilePrefix, data.Id)
	gomallUserMobileKey := fmt.Sprintf("%s%v", cacheGomallUserUserMobilePrefix, data.Mobile)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info)
		}
		return conn.ExecCtx(ctx, query, data.DelState, data.Version, data.Mobile, data.Password, data.Nickname, data.Sex, data.Avatar, data.Info)
	}, gomallUserIdKey, gomallUserMobileKey)
}
