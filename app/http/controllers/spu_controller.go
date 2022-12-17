package controllers

import (
	"quick-go/app/entity"
	"quick-go/app/service"
	"quick-go/global/consts"
	"quick-go/utils/response"

	"github.com/gin-gonic/gin"
)

type SpuController struct {
	SpuService service.SpuService
}

func NewSpuController(ss service.SpuService) SpuController {
	controller := SpuController{
		SpuService: ss,
	}
	return controller
}

// GetSpuInfo 获取商品信息
func (sc *SpuController) GetSpuInfo(c *gin.Context) {
	// 参数校验
	req := entity.GetSpuInfoReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, err.Error(), err)
		return
	}

	// 调用service
	data, err := sc.SpuService.GetSpuInfo(&req)
	response.Respond(c, data, err)
}

// CreateSpu create a new spu
func (sc *SpuController) CreateSpu(c *gin.Context) {
	// 参数校验
	req := entity.CreateSpuReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, err.Error(), err)
		return
	}

	// data := req
	// var err error

	// 调用service
	data, err := sc.SpuService.CreateSpu(&req)
	response.Respond(c, data, err)
}
func (sc *SpuController) AsyncRedisList(c *gin.Context) {
	// 参数校验
	req := entity.GetSpuInfoReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, err.Error(), err)
		return
	}

	// 调用service
	data, err := sc.SpuService.AsyncRedisList(&req)
	response.Respond(c, data, err)
}
