// Code generated by goctl. DO NOT EDIT.
package types

type CreateSeckillReq struct {
	SeckillId       int64 `json:"seckillId"`
	ProductId       int64 `json:"productId"`
	ProductQuantity int64 `json:"productQuantity"`
}

type CreatSeckillResp struct {
	SeckillSn string `json:"seckillSn"`
}
