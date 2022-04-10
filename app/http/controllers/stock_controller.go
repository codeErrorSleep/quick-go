package controllers

import (
	"github.com/gin-gonic/gin"
	"quick-go/app/entity"
	"quick-go/app/service"
	consts "quick-go/global"
	"quick-go/utils/response"
)

// ReduceStock 删减库存
func ReduceStock(c *gin.Context) {

}

// GetSpuStock 获取商品库存
func GetSpuStock(c *gin.Context) {
	// 参数校验
	req := entity.GetSpuStockReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, consts.ValidatorParamsCheckFailMsg, err)
		return
	}

	// 调用service
	svc := service.StockServiceNew(c)
	resData, err := svc.GetSpuStock(&req)
	if err != nil {
		response.Fail(c, consts.CurdSelectErrorCode, consts.CurdSelectErrorMsg, err)
	}
	response.Success(c, resData)
}
