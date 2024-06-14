package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PayinfoModel = (*customPayinfoModel)(nil)

type (
	// PayinfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPayinfoModel.
	PayinfoModel interface {
		payinfoModel
	}

	customPayinfoModel struct {
		*defaultPayinfoModel
	}
)

// NewPayinfoModel returns a model for the database table.
func NewPayinfoModel(conn sqlx.SqlConn, c cache.CacheConf) PayinfoModel {
	return &customPayinfoModel{
		defaultPayinfoModel: newPayinfoModel(conn, c),
	}
}
