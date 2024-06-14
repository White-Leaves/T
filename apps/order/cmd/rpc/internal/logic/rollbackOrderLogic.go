package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gomall/apps/product/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gomall/apps/order/cmd/rpc/internal/svc"
	"gomall/apps/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackOrderLogic {
	return &RollbackOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RollbackOrderLogic) RollbackOrder(in *pb.CreateOrderReq) (*pb.CreateOrderResp, error) {
	fmt.Printf("RollbackOrder:%+v\n", in)

	order, err := l.svcCtx.OrderModel.FindOneByUid(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if order != nil {

		barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
		db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
		if err != nil {
			return nil,
				status.Error(codes.Internal, err.Error())
		}
		if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
			order.Status = 60

			if err := l.svcCtx.OrderModel.TxUpdate(tx, order); err != nil {
				return fmt.Errorf("订单回滚失败")
			}
			return nil
		}); err != nil {
			logx.Errorf("err:%v\n", err)
		}
	}

	return &pb.CreateOrderResp{}, nil
}
