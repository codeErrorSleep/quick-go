package entity

type AppRequest struct {
	AppID string `json:"app_id"`
	SpuID string `json:"spu_id"`
}

type DefaultRequest struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
