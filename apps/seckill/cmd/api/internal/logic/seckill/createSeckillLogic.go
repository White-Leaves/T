package seckill

import (
	"context"

	"gomall/apps/seckill/cmd/api/internal/svc"
	"gomall/apps/seckill/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSeckillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillLogic {
	return &CreateSeckillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSeckillLogic) CreateSeckill(req *types.CreateSeckillReq) (resp *types.CreatSeckillResp, err error) {

	return
}
