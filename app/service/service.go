package service

import "quick-go/app/entity"

type IStockService interface {
	GetSpuStock(req *entity.GetSpuStockReq) (resData *entity.GetSpuStockRes, err error)
}

type ISpuService interface {
	GetSpuInfo(req *entity.GetSpuInfoReq) (res *entity.GetSpuInfoRes, err error)
	AsyncRedisList(req *entity.GetSpuInfoReq) (res *entity.GetSpuInfoRes, err error)
}
