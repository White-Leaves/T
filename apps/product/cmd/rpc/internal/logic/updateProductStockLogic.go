package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gomall/apps/product/cmd/rpc/internal/svc"
	"gomall/apps/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStockLogic {
	return &UpdateProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductStockLogic) UpdateProductStock(in *pb.UpdateProductStockReq) (*pb.UpdateProductStockResp, error) {
	err := l.svcCtx.ProductModel.UpdateProductStock(l.ctx, in.ProductId, in.Num)
	if err != nil {
		//重试
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.UpdateProductStockResp{}, nil
}
