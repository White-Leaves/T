package logic

import (
	"context"
	"github.com/dtm-labs/dtmcli"
	"gomall/apps/product/cmd/rpc/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
	"gomall/apps/product/cmd/rpc/internal/svc"
)

type CheckProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckProductStockLogic {
	return &CheckProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckProductStockLogic) CheckProductStock(in *product.UpdateProductStockReq) (*product.UpdateProductStockResp, error) {
	p, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.ProductId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if p.Stock < in.Num {
		return nil, status.Error(codes.ResourceExhausted, dtmcli.ResultFailure)
	}
	return &product.UpdateProductStockResp{}, nil
}
