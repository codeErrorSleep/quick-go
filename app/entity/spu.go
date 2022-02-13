package entity

type GetSpuInfo struct {
	AppId string `json:"app_id" binding:"required"`
	SpuId string `json:"spu_id" binding:"required"`
	SkuId string `json:"sku_id" binding:"required"`
}
