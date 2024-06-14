package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gomall/apps/user/model"

	"gomall/apps/product/cmd/rpc/internal/svc"
	"gomall/apps/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductDetailLogic) ProductDetail(in *pb.ProductDetailReq) (*pb.ProductDetailResp, error) {
	product, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)

	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("Product DB error" + err.Error())
	}

	var pbProduct pb.ProductItem
	if product != nil {
		_ = copier.Copy(&pbProduct, product)
	}

	return &pb.ProductDetailResp{
		Product: &pbProduct,
	}, nil
}
