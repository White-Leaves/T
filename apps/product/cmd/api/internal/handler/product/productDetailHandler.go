package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gomall/apps/product/cmd/api/internal/logic/product"
	"gomall/apps/product/cmd/api/internal/svc"
	"gomall/apps/product/cmd/api/internal/types"
)

func ProductDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewProductDetailLogic(r.Context(), svcCtx)
		resp, err := l.ProductDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
