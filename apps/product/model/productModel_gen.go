// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	productFieldNames          = builder.RawFieldNames(&Product{})
	productRows                = strings.Join(productFieldNames, ",")
	productRowsExpectAutoSet   = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	productRowsWithPlaceHolder = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGomallProductProductIdPrefix = "cache:gomallProduct:product:id:"
)

type (
	productModel interface {
		Insert(ctx context.Context, data *Product) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Product, error)
		Update(ctx context.Context, data *Product) error
		Delete(ctx context.Context, id int64) error
	}

	defaultProductModel struct {
		sqlc.CachedConn
		table string
	}

	Product struct {
		Id           int64     `db:"id"`
		CreateTime   time.Time `db:"create_time"`
		UpdateTime   time.Time `db:"update_time"`
		DelState     int64     `db:"del_state"`
		Name         string    `db:"name"`          // 名称
		Info         string    `db:"info"`          // 介绍
		Price        int64     `db:"price"`         // 价格
		Stock        int64     `db:"stock"`         // 库存数量
		Status       int64     `db:"status"`        // 商品状态1.在售2.下架3.删除
		SeckillStock int64     `db:"seckill_stock"` // 秒杀库存
	}
)

func newProductModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultProductModel {
	return &defaultProductModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`product`",
	}
}

func (m *defaultProductModel) Delete(ctx context.Context, id int64) error {
	gomallProductProductIdKey := fmt.Sprintf("%s%v", cacheGomallProductProductIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, gomallProductProductIdKey)
	return err
}

func (m *defaultProductModel) FindOne(ctx context.Context, id int64) (*Product, error) {
	gomallProductProductIdKey := fmt.Sprintf("%s%v", cacheGomallProductProductIdPrefix, id)
	var resp Product
	err := m.QueryRowCtx(ctx, &resp, gomallProductProductIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductModel) Insert(ctx context.Context, data *Product) (sql.Result, error) {
	gomallProductProductIdKey := fmt.Sprintf("%s%v", cacheGomallProductProductIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, productRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DelState, data.Name, data.Info, data.Price, data.Stock, data.Status, data.SeckillStock)
	}, gomallProductProductIdKey)
	return ret, err
}

func (m *defaultProductModel) Update(ctx context.Context, data *Product) error {
	gomallProductProductIdKey := fmt.Sprintf("%s%v", cacheGomallProductProductIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DelState, data.Name, data.Info, data.Price, data.Stock, data.Status, data.SeckillStock, data.Id)
	}, gomallProductProductIdKey)
	return err
}

func (m *defaultProductModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGomallProductProductIdPrefix, primary)
}

func (m *defaultProductModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultProductModel) tableName() string {
	return m.table
}