package service

import (
	"context"
	"fmt"
	"quick-go/app/entity"
	"quick-go/app/repo"
	"quick-go/global/consts"
	"quick-go/utils/quickErrors"
)

type StockService struct {
	stockRepo repo.IStockRepo
}

func NewStockService(stockRepo repo.IStockRepo) IStockService {
	svc := StockService{stockRepo: stockRepo}
	return &svc
}

// getSpuStock 获取商品的信息
func (s *StockService) GetSpuStock(req *entity.GetSpuStockReq) (resData *entity.GetSpuStockRes, err error) {
	appID := req.AppID
	spuID := req.SpuID

	// 获取stock的信息
	stockList, err := s.stockRepo.GetStockDetail(context.TODO(), appID, spuID)
	if err != nil {
		return nil, quickErrors.New(consts.CurdSelectFailCode, fmt.Sprint(req.AppID, req.SpuID), consts.CurdSelectFailMsg)
	}

	// 组装返回参数
	resData = &entity.GetSpuStockRes{
		AppID: appID,
		SpuID: spuID,
	}
	for _, stockInfo := range stockList {
		resData.SkuInfo = append(resData.SkuInfo, entity.SkuInfoRes{
			SkuID:   stockInfo.SkuID,
			SellNum: stockInfo.SellNum,
			LeftNum: stockInfo.LeftNum,
		})
	}

	return resData, nil
}
