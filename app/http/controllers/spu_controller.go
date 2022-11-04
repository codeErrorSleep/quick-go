package controllers

import (
	"quick-go/app/entity"
	"quick-go/app/service"
	"quick-go/global/consts"
	"quick-go/utils/response"

	"github.com/gin-gonic/gin"
)

// GetSpuInfo 获取商品信息
func GetSpuInfo(c *gin.Context) {
	// 参数校验
	req := entity.GetSpuInfoReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, err.Error(), err)
		return
	}

	// 调用service
	svc := service.SpuServiceNew(c)
	data, err := svc.GetSpuInfo(&req)
	response.Respond(c, data, err)
}

// CreateSpu create a new spu
func CreateSpu(c *gin.Context) {
	// 参数校验
	req := entity.CreateSpuReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, err.Error(), err)
		return
	}

	data := req
	var err error

	// 调用service
	// svc := service.SpuServiceNew(c)
	// data, err := svc.CreateSpu(&req)
	response.Respond(c, data, err)
}
func AsyncRedisList(c *gin.Context) {
	// 参数校验
	req := entity.GetSpuInfoReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, err.Error(), err)
		return
	}

	// 调用service
	svc := service.SpuServiceNew(c)
	data, err := svc.AsyncRedisList(&req)
	response.Respond(c, data, err)
}
