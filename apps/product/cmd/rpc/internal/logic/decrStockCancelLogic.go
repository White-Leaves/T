package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gomall/apps/product/cmd/rpc/internal/svc"
	"gomall/apps/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrStockCancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockCancelLogic {
	return &DecrStockCancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockCancelLogic) DecrStockCancel(in *pb.DecrStockReq) (*pb.DecrStockResp, error) {
	err := l.svcCtx.ProductModel.UpdateProductStock(l.ctx, in.Id, -in.Num)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DecrStockResp{}, nil
}
