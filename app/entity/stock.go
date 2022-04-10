package entity

type GetSpuStockReq struct {
	AppID string `json:"app_id" binding:"required"`
	SpuID string `json:"spu_id" binding:"required"`
}

type GetSpuStockRes struct {
	AppID   string       `json:"app_id"`
	SpuID   string       `json:"spu_id"`
	SkuInfo []SkuInfoRes `json:"sku_info"`
}
type SkuInfoRes struct {
	SkuID   string `json:"sku_id"`
	SellNum int64  `json:"sell_num"`
	LeftNum int64  `json:"left_num"`
}
