package entity

type GetSpuInfo struct {
	AppID string `json:"app_id" binding:"required"`
	SpuID string `json:"spu_id" binding:"required"`
	SkuID string `json:"sku_id" binding:"required"`
}
