package entity

type CreateSpuReq struct {
	AppID     string `json:"app_id" binding:"required"`
	GoodsName string `json:"goods_name" binding:"required"`
	Status    int    `json:"status" binding:"min=0"`
	Price     int    `json:"price" binding:"min=0"`
	PriceLine int    `json:"price_line" binding:"min=0"`
}

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
	SaleAt       int64  `json:"sale_at"`
}
