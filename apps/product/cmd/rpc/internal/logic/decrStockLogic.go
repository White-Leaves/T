package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gomall/apps/product/cmd/rpc/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
	"gomall/apps/product/cmd/rpc/internal/svc"
)

type DecrStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockLogic {
	return &DecrStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockLogic) DecrStock(in *product.DecrStockReq) (*product.DecrStockResp, error) {
	stock, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if stock == nil || stock.Stock < in.Num {
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 更新产品库存
		result, err := l.svcCtx.ProductModel.TxUpdateStock(tx, in.Id, int(in.Num))
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		affected, err := result.RowsAffected()
		// 库存不足，返回子事务失败
		if err == nil && affected <= 0 {
			return status.Error(codes.Aborted, dtmcli.ResultFailure)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &product.DecrStockResp{}, nil

}
