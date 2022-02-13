package entity

type GetSpuStock struct {
	AppId string `json:"app_id" binding:"required"`
	SpuId string `json:"spu_id" binding:"required"`
}
