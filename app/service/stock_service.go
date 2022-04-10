package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"quick-go/app/entity"
	"quick-go/db"
	"quick-go/db/models"
)

type StockService struct {
	ctx *gin.Context
}

func StockServiceNew(ctx *gin.Context) *StockService {
	svc := StockService{ctx: ctx}
	return &svc
}

// getSpuStock 获取商品的信息
func (s *StockService) GetSpuStock(req *entity.GetSpuStockReq) (resData entity.GetSpuStockRes, err error) {
	appID := req.AppID
	spuID := req.SpuID
	stockRedisKey := appID + ":" + spuID
	// 查redis
	stockRedisDetail := db.RedisLocal.HGetAll(stockRedisKey)
	if stockRedisDetail.Err() != nil {
		return resData, stockRedisDetail.Err()
	}
	if stockRedisDetail.Val() != nil {
		fmt.Println(stockRedisDetail.Val())
	}

	// 获取stock的信息
	stock := models.Stock{}
	stockList, err := stock.GetStockDetail(appID, spuID)
	if err != nil {
		return resData, err
	}

	// 存到redis
	stockRedisMap := make(map[string]interface{})
	for i := 0; i < len(stockList); i++ {
		stockRedisMap[stockList[i].SkuID] = stockList[i].LeftNum
	}
	statusCmd := db.RedisLocal.HMSet(stockRedisKey, stockRedisMap)
	if statusCmd.Err() != nil {
		return resData, stockRedisDetail.Err()
	}

	// 组装返回参数
	resData = entity.GetSpuStockRes{
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
