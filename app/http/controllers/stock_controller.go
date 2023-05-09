package controllers

import (
	"quick-go/app/entity"
	"quick-go/app/service"
	"quick-go/global/consts"
	"quick-go/utils/response"

	"github.com/gin-gonic/gin"
)

type StockController struct {
	StockService service.IStockService
}

func NewStockController(ss service.IStockService) StockController {
	controller := StockController{
		StockService: ss,
	}
	return controller
}

// ReduceStock 删减库存
func (sc *StockController) ReduceStock(c *gin.Context) {

}

// GetSpuStock 获取商品库存
func (sc *StockController) GetSpuStock(c *gin.Context) {
	// 参数校验
	req := entity.GetSpuStockReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, err)
		return
	}

	// // 调用service
	// svc := service.StockServiceNew(c)
	// resData, err := svc.GetSpuStock(&req)
	resData, err := sc.StockService.GetSpuStock(&req)

	// resData := map[string]interface{}{}
	// var err error
	response.Respond(c, resData, err)
}

// ReduceStock 删减库存
func (sc *StockController) ReduceStock1(c *gin.Context) {

}

// ReduceStock 删减库存
func (sc *StockController) ReduceStock2(c *gin.Context) {

}

// ReduceStock 删减库存
func (sc *StockController) ReduceStock3(c *gin.Context) {

}

// ReduceStock 删减库存

// ReduceStock 删减库存
// ReduceStock 删减库存

// ReduceStock 删减库存
// ReduceStock 删减库存

// ReduceStock 删减库存
// ReduceStock 删减库存

// ReduceStock 删减库存
// ReduceStock 删减库存

// ReduceStock 删减库存
