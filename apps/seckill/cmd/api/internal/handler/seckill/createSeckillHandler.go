package seckill

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gomall/apps/seckill/cmd/api/internal/logic/seckill"
	"gomall/apps/seckill/cmd/api/internal/svc"
	"gomall/apps/seckill/cmd/api/internal/types"
)

func CreateSeckillHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSeckillReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := seckill.NewCreateSeckillLogic(r.Context(), svcCtx)
		resp, err := l.CreateSeckill(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
