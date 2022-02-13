package service

import (
	"github.com/gin-gonic/gin"
	"quick-go/app/entity"
)

type SpuService struct {
	ctx *gin.Context
}

func SpuServiceNew(ctx *gin.Context) *SpuService {
	svc := SpuService{ctx: ctx}
	return &svc
}

// getSpuInfo 获取商品的信息
func (s *SpuService) getSpuInfo(req *entity.GetSpuInfo) {
	// 查redis

	// 查mysql

	// 存到redis
}
