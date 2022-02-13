package controllers

import (
	"github.com/gin-gonic/gin"
	"quick-go/app/entity"
	consts "quick-go/global"
	"quick-go/utils/response"
)

// GetSpuInfo 获取商品信息
func GetSpuInfo(c *gin.Context) {
	// 参数校验
	req := entity.GetSpuInfo{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, err.Error(), err)
		return
	}

	// 调用service
	//svc := service.SpuServiceNew(c)
	//data, err := svc.GetDeliveryPlaceList(&req)
}
