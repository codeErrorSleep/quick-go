package service

import "quick-go/app/entity"

type IStockService interface {
	GetSpuStock(req *entity.GetSpuStockReq) (resData *entity.GetSpuStockRes, err error)
}
