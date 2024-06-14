package logic

import (
	"context"
	"encoding/json"
	"gomall/apps/order/cmd/api/internal/svc"
	"gomall/apps/order/cmd/api/internal/types"
	"gomall/apps/order/cmd/rpc/order"
	"google.golang.org/grpc/codes"

	"gomall/apps/product/cmd/rpc/product"

	"google.golang.org/grpc/status"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.AddOrderReq) (*types.AddOrderResp, error) {
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, err
	}

	orderRpcServer, err := l.svcCtx.Config.OrderRpcConf.BuildTarget()
	if err != nil {
		return nil, err
	}

	productRpcServer, err := l.svcCtx.Config.ProductRpcConf.BuildTarget()
	if err != nil {
		return nil, err
	}

	var dtmServer = "etcd://localhost:23789chuang/dtmservice"

	gid := dtmgrpc.MustGenGid(dtmServer)

	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcServer+"/order.Order/CreateOrderDTM",
			orderRpcServer+"/order.Order/rollbackOrder",
			&order.AddOrderReq{
				Userid:    uid,
				Productid: req.ProductId,
				Quantity:  req.Quantity,
				Postage:   req.Postage,
			}).Add(productRpcServer+"/product.Product/DecrStock",
		productRpcServer+"/product.Product/rollbackProductStock",
		&product.DecrStockReq{
			Id:  req.ProductId,
			Num: req.Quantity,
		})

	err = saga.Submit()
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	return &types.AddOrderResp{
		Sn: "test",
	}, nil
	//
	//
	//	ProductResp, err := l.svcCtx.ProductRpc.ProductDetail(l.ctx, &pb.ProductDetailReq{
	//		Id: req.ProductId,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	if ProductResp.Product == nil || ProductResp.Product.Id == 0 {
	//		return nil, errors.New("product not found")
	//
	//	}
	//
	//	userId := ctxdata.GetUidFromCtx(l.ctx)
	//
	//	resp, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &order.CreateOrderReq{
	//
	//		UserId:    userId,
	//		ProductId: req.ProductId,
	//		Quantity:  req.Quantity,
	//	})
	//	if err != nil {
	//		return nil, err
	//	}

	//return &types.CreatOrderResp{
	//	OrderSn: resp.Sn,
	//}, nil

}
