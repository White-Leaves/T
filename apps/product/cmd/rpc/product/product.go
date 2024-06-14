// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package product

import (
	"context"

	"gomall/apps/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckAndUpdateStockReq  = pb.CheckAndUpdateStockReq
	CheckAndUpdateStockResp = pb.CheckAndUpdateStockResp
	DecrStockReq            = pb.DecrStockReq
	DecrStockResp           = pb.DecrStockResp
	ProductDetailReq        = pb.ProductDetailReq
	ProductDetailResp       = pb.ProductDetailResp
	ProductItem             = pb.ProductItem
	ProductItemRequest      = pb.ProductItemRequest
	UpdateProductStockReq   = pb.UpdateProductStockReq
	UpdateProductStockResp  = pb.UpdateProductStockResp

	Product interface {
		Product(ctx context.Context, in *ProductItemRequest, opts ...grpc.CallOption) (*ProductItem, error)
		ProductDetail(ctx context.Context, in *ProductDetailReq, opts ...grpc.CallOption) (*ProductDetailResp, error)
		UpdateProductStock(ctx context.Context, in *UpdateProductStockReq, opts ...grpc.CallOption) (*UpdateProductStockResp, error)
		CheckAndUpdateStock(ctx context.Context, in *CheckAndUpdateStockReq, opts ...grpc.CallOption) (*CheckAndUpdateStockResp, error)
		CheckProductStock(ctx context.Context, in *UpdateProductStockReq, opts ...grpc.CallOption) (*UpdateProductStockResp, error)
		RollbackProductStock(ctx context.Context, in *UpdateProductStockReq, opts ...grpc.CallOption) (*UpdateProductStockResp, error)
		DecrStock(ctx context.Context, in *DecrStockReq, opts ...grpc.CallOption) (*DecrStockResp, error)
		DecrStockCancel(ctx context.Context, in *DecrStockReq, opts ...grpc.CallOption) (*DecrStockResp, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Product(ctx context.Context, in *ProductItemRequest, opts ...grpc.CallOption) (*ProductItem, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Product(ctx, in, opts...)
}

func (m *defaultProduct) ProductDetail(ctx context.Context, in *ProductDetailReq, opts ...grpc.CallOption) (*ProductDetailResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductDetail(ctx, in, opts...)
}

func (m *defaultProduct) UpdateProductStock(ctx context.Context, in *UpdateProductStockReq, opts ...grpc.CallOption) (*UpdateProductStockResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.UpdateProductStock(ctx, in, opts...)
}

func (m *defaultProduct) CheckAndUpdateStock(ctx context.Context, in *CheckAndUpdateStockReq, opts ...grpc.CallOption) (*CheckAndUpdateStockResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CheckAndUpdateStock(ctx, in, opts...)
}

func (m *defaultProduct) CheckProductStock(ctx context.Context, in *UpdateProductStockReq, opts ...grpc.CallOption) (*UpdateProductStockResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CheckProductStock(ctx, in, opts...)
}

func (m *defaultProduct) RollbackProductStock(ctx context.Context, in *UpdateProductStockReq, opts ...grpc.CallOption) (*UpdateProductStockResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.RollbackProductStock(ctx, in, opts...)
}

func (m *defaultProduct) DecrStock(ctx context.Context, in *DecrStockReq, opts ...grpc.CallOption) (*DecrStockResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.DecrStock(ctx, in, opts...)
}

func (m *defaultProduct) DecrStockCancel(ctx context.Context, in *DecrStockReq, opts ...grpc.CallOption) (*DecrStockResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.DecrStockCancel(ctx, in, opts...)
}
