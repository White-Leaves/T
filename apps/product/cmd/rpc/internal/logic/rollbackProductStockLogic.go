package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gomall/apps/product/cmd/rpc/internal/svc"
	"gomall/apps/product/cmd/rpc/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RollbackProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackProductStockLogic {
	return &RollbackProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RollbackProductStockLogic) RollbackProductStock(in *product.UpdateProductStockReq) (*product.UpdateProductStockResp, error) {

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		if _, err := l.svcCtx.ProductModel.TxUpdateStock(tx, in.ProductId, int(-in.Num)); err != nil {
			return fmt.Errorf("库存回滚失败", in.ProductId, in.Num)
		}
		return nil
	}); err != nil {
		logx.Errorf("err:%v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &product.UpdateProductStockResp{}, nil

}
