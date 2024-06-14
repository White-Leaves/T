package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gomall/apps/order/cmd/api/internal/logic"
	"gomall/apps/order/cmd/api/internal/svc"
	"gomall/apps/order/cmd/api/internal/types"
)

func OrderDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderDetailLogic(r.Context(), svcCtx)
		resp, err := l.OrderDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
