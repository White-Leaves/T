package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gomall/apps/order/model"
	"gomall/apps/product/cmd/rpc/product"
	"gomall/apps/user/cmd/rpc/user"
	"gomall/pkg/snowflake"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
	"gomall/apps/order/cmd/rpc/internal/svc"
	"gomall/apps/order/cmd/rpc/pb"
)

type CreateOrderDTMLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderDTMLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderDTMLogic {
	return &CreateOrderDTMLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderDTMLogic) CreateOrderDTM(in *pb.AddOrderReq) (*pb.AddOrderResp, error) {
	var (
		userRpcResp    *user.GetUserInfoResp
		productRpcResp *product.ProductItem
	)

	checkProduct := func() error {
		var err error
		var productReq product.ProductItemRequest
		productReq.Id = in.Productid
		productRpcResp, err = l.svcCtx.ProductRpc.Product(l.ctx, &productReq)
		if err != nil {
			return nil
		}
		return nil
	}
	checkUser := func() error {
		var err error
		var userReq user.GetUserInfoReq
		userReq.Id = in.Userid
		userRpcResp, err = l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userReq)
		if err != nil {
			return nil
		}
		return nil
	}

	err := mr.Finish(checkProduct, checkUser)

	if userRpcResp == nil {
		return nil, errors.New("User not exist")
	}
	if productRpcResp == nil {
		return nil, errors.New("product not exist")
	}

	if productRpcResp.Stock <= in.Quantity {
		return nil, errors.New("product stock not enough")
	}

	orderSn := snowflake.GenString()

	db, err := sqlx.NewMysql(l.svcCtx.Config.DB.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		insertOrderitem := model.Orderitem{
			Sn:           orderSn,
			UserId:       in.Userid,
			ProductId:    in.Productid,
			ProductName:  productRpcResp.Name,
			CurrentPrice: float64(productRpcResp.Price),
			Quantity:     in.Quantity,
			TotalPrice:   float64(in.Quantity * productRpcResp.GetPrice()),
		}

		insertOrderitemResp, err := l.svcCtx.OrderitemModel.Insert(l.ctx, &insertOrderitem)
		if err != nil {
			return nil
		}
		_, err = insertOrderitemResp.LastInsertId()
		if err != nil {
			return err
		}

		insertOrder := model.Orders{
			Sn:      orderSn,
			UserId:  in.Userid,
			Postage: float64(in.Postage),
		}

		_, err = l.svcCtx.OrderModel.TxInsert(tx, &insertOrder)
		if err != nil {
			return nil
		}
		return nil

	}); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.AddOrderResp{
		Sn: orderSn,
	}, nil
}
