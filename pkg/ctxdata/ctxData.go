package ctxdata

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

//CtxKeyJwtUserId

var CtxKeyJwtUserId = "JwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64uid, err := jsonUid.Int64(); err == nil {
			uid = int64uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err ", err)

		}

	}

	return uid
}
