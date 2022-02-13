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
func (s *StockService) GetSpuStock(req *entity.GetSpuStock) error {
	appId := req.AppId
	spuId := req.SpuId
	stockRedisKey := appId + ":" + spuId
	// 查redis
	stockRedisDetail := db.RedisLocal.HGetAll(stockRedisKey)
	if stockRedisDetail.Err() != nil {
		return stockRedisDetail.Err()
	}
	if stockRedisDetail.Val() != nil {
		fmt.Println(stockRedisDetail.Val())
	}

	// 获取stock的信息
	stock := models.Stock{}
	stockList, err := stock.GetStockDetail(appId, spuId)
	if err != nil {
		return err
	}

	// 存到redis
	stockRedisMap := make(map[string]interface{})
	for i := 0; i < len(stockList); i++ {
		stockRedisMap[stockList[i].SkuID] = stockList[i].LeftNum
	}

	statusCmd := db.RedisLocal.HMSet(stockRedisKey, stockRedisMap)
	if statusCmd.Err() != nil {
		return stockRedisDetail.Err()
	}

	return nil
}
