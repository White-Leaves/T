package logic

import (
	"context"
	"fmt"
	"gomall/apps/product/cmd/rpc/product"
	"gomall/apps/product/model"

	"gomall/apps/product/cmd/rpc/internal/svc"
	"gomall/apps/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLogic) Product(in *pb.ProductItemRequest) (*pb.ProductItem, error) {
	v, err, _ := l.svcCtx.SingleGroup.Do(fmt.Sprintf("product:%d", in.Id), func() (interface{}, error) {
		return l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	})
	if err != nil {
		return nil, err
	}

	p := v.(*model.Product)
	return &product.ProductItem{
		Id:    p.Id,
		Name:  p.Name,
		Stock: p.Stock,
	}, nil

}
