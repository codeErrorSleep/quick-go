package service

import (
	"context"
	"quick-go/app/rpc"

	"github.com/gin-gonic/gin"
)

// test rpc
type HelloService struct {
	// ctx *gin.Context
	rpc.UnimplementedHelloServiceServer
}

func HelloServiceNew(ctx *gin.Context) *HelloService {
	return &HelloService{
		// ctx: ctx,
	}
}

func (hs *HelloService) GetBrandInfo(ctx context.Context, req *rpc.GetBrandInfoReq) (*rpc.GetBrandInfoRsp, error) {
	if req.BrandId == 10 {
		return &rpc.GetBrandInfoRsp{BrandId: 10, BrandName: "ad"}, nil
	}
	return &rpc.GetBrandInfoRsp{BrandId: req.BrandId, BrandName: "other"}, nil
}
