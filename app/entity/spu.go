package entity

type GetSpuInfoReq struct {
	AppID string `json:"app_id" binding:"required"`
	SpuID string `json:"spu_id" binding:"required"`
	SkuID string `json:"sku_id" binding:"required"`
}

type GetSpuInfoRes struct {
	AppID        string `json:"app_id"`
	SpuID        string `json:"spu_id"`
	SpuType      string `json:"spu_type"`
	ResourceID   string `json:"resource_id"`
	ResourceType int    `json:"resource_type"`
	GoodsName    string `json:"goods_name"`
	GoodsImg     string `json:"goods_img"`
}
