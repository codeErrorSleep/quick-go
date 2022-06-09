package service

import (
	"fmt"
	"quick-go/app/entity"
	"quick-go/app/models"
	"quick-go/global/consts"
	"quick-go/utils/errors"

	"github.com/gin-gonic/gin"
)

type StockService struct {
	ctx *gin.Context
}

func StockServiceNew(ctx *gin.Context) *StockService {
	svc := StockService{ctx: ctx}
	return &svc
}

// getSpuStock 获取商品的信息
func (s *StockService) GetSpuStock(req *entity.GetSpuStockReq) (resData *entity.GetSpuStockRes, err error) {
	appID := req.AppID
	spuID := req.SpuID

	// 获取stock的信息
	stock := models.Stock{}
	stockList, err := stock.GetStockDetail(appID, spuID)
	if err != nil {
		return nil, errors.New(consts.CurdSelectFailCode, fmt.Sprint(req.AppID, req.SpuID), consts.CurdSelectFailMsg)
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
