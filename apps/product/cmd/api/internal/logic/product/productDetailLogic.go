package product

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gomall/apps/product/cmd/rpc/product"

	"gomall/apps/product/cmd/api/internal/svc"
	"gomall/apps/product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDetailLogic) ProductDetail(req *types.ProductDetailReq) (resp *types.ProductDetailResp, err error) {

	productResp, err := l.svcCtx.ProductRpc.ProductDetail(l.ctx, &product.ProductDetailReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, errors.New("get product detail fail,err:" + err.Error())
	}

	var typeProduct types.Product
	if productResp.Product != nil {

		_ = copier.Copy(&typeProduct, productResp.Product)
	}

	return &types.ProductDetailResp{
		Product: typeProduct,
	}, nil

}
