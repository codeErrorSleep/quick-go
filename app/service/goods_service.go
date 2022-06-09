package service

import (
	"quick-go/app/entity"
	"quick-go/app/repo"

	"github.com/gin-gonic/gin"
)

type SpuService struct {
	ctx     *gin.Context
	SpuRepo repo.ISpuRepo
}

func SpuServiceNew(ctx *gin.Context) *SpuService {
	svc := SpuService{ctx: ctx, SpuRepo: repo.NewMysqlSpuRepository()}
	return &svc
}

// getSpuInfo 获取商品的信息
func (s *SpuService) getSpuInfo(req *entity.GetSpuInfo) {
}
