package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"quick-go/app/entity"
	"quick-go/db"
	"quick-go/db/models"
	"time"
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
	stockRedisKey := "stock-" + appID + ":" + spuID
	// 查redis
	stockRedisDetail := db.RedisLocal.Get(stockRedisKey)
	if stockRedisDetail.Val() != "" {
		stockRedisByte, _ := stockRedisDetail.Bytes()
		err = json.Unmarshal(stockRedisByte, &resData)
		if err != nil {
			return resData, err
		}
		return resData, nil
	}

	// 获取stock的信息
	stock := models.Stock{}
	stockList, err := stock.GetStockDetail(appID, spuID)
	if err != nil {
		return resData, err
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

	// 设置redis
	stockStr, err := json.Marshal(resData)
	if err != nil {
		return resData, err
	}
	db.RedisLocal.Set(stockRedisKey, stockStr, time.Second*60)

	return resData, nil
}
