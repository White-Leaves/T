package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gomall/apps/order/cmd/api/internal/logic"
	"gomall/apps/order/cmd/api/internal/svc"
	"gomall/apps/order/cmd/api/internal/types"
)

func createOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateOrderLogic(r.Context(), svcCtx)
		resp, err := l.CreateOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
