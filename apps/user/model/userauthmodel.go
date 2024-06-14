package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAuthModel = (*customUserAuthModel)(nil)

type (
	// UserAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthModel.
	UserAuthModel interface {
		userAuthModel
		InsertR(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error)
	}

	customUserAuthModel struct {
		*defaultUserAuthModel
	}
)

// NewUserAuthModel returns a model for the database table.
func NewUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf) UserAuthModel {
	return &customUserAuthModel{
		defaultUserAuthModel: newUserAuthModel(conn, c),
	}
}

func (m *defaultUserAuthModel) InsertR(ctx context.Context, session sqlx.Session, data *UserAuth) (sql.Result, error) {
	gomallUserAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheGomallUserUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	gomallUserAuthIdKey := fmt.Sprintf("%s%v", cacheGomallUserUserAuthIdPrefix, data.Id)
	gomallUserAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheGomallUserUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, userAuthRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.DelState, data.Version, data.UserId, data.AuthKey, data.AuthType)
		}
		return conn.ExecCtx(ctx, query, data.DelState, data.Version, data.UserId, data.AuthKey, data.AuthType)
	}, gomallUserAuthAuthTypeAuthKeyKey, gomallUserAuthIdKey, gomallUserAuthUserIdAuthTypeKey)
}
