package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"gomall/apps/order/cmd/rpc/internal/svc"
	"gomall/apps/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderBySnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderBySnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderBySnLogic {
	return &GetOrderBySnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderBySnLogic) GetOrderBySn(in *pb.GetOrderBySnReq) (*pb.GetOrderBySnResp, error) {

	Order, err := l.svcCtx.OrderModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil {
		return nil, errors.Wrapf(err, "HomestayOrderModel  FindOneBySn db err : %v , sn : %s", err, in.Sn)
	}

	var resp pb.Orders

	if Order != nil {
		_ = copier.Copy(&resp, Order)

		resp.CreateTime = Order.CreateTime.Unix()

	}

	return &pb.GetOrderBySnResp{
		Order: &resp,
	}, nil
}
