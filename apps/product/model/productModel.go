package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProductModel = (*customProductModel)(nil)

type (
	// ProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProductModel.
	ProductModel interface {
		productModel
		UpdateProductStock(ctx context.Context, pid, num int64) error
		TxUpdateStock(tx *sql.Tx, id int64, delta int) (sql.Result, error)
	}

	customProductModel struct {
		*defaultProductModel
	}
)

// NewProductModel returns a model for the database table.
func NewProductModel(conn sqlx.SqlConn, c cache.CacheConf) ProductModel {
	return &customProductModel{
		defaultProductModel: newProductModel(conn, c),
	}
}

func (m *customProductModel) UpdateProductStock(ctx context.Context, pid, num int64) error {
	productIdKey := fmt.Sprintf("%s%v", cacheGomallProductProductIdPrefix, pid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("UPDATE %s SET stock = stock - ? WHERE id = ? and stock > 0", m.table), num, pid)

	}, productIdKey)
	return err
}

func (m *customProductModel) UpdateProductSeckillStock(ctx context.Context, pid, num int64) error {
	productIdKey := fmt.Sprintf("%s%v", cacheGomallProductProductIdPrefix, pid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, fmt.Sprintf("UPDATE %s SET seckill_stock = seckill_stock - ? WHERE id = ? and stock > 0", m.table), num, pid)

	}, productIdKey)
	return err
}

func (m *customProductModel) TxUpdateStock(tx *sql.Tx, id int64, delta int) (sql.Result, error) {
	productIdKey := fmt.Sprintf("%s%v", cacheGomallProductProductIdPrefix, id)
	return m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set stock = stock + ? where stock >= -? and id=?", m.table)
		return tx.Exec(query, delta, delta, id)
	}, productIdKey)
}
