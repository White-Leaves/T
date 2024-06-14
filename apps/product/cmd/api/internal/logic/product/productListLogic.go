package product

import (
	"context"
	"fmt"
	"gomall/apps/product/cmd/api/internal/svc"
	"gomall/apps/product/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ProductListReq) (resp *types.ProductListResp, err error) {
	fmt.Println("product list")

	return
}
