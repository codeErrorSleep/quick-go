package service

import (
	"context"
	"encoding/json"
	"fmt"
	"quick-go/app/entity"
	"quick-go/app/repo"
	"quick-go/global"
	"quick-go/global/consts"
	"quick-go/utils/quickErrors"
	"time"

	"github.com/Shopify/sarama"
)

type SpuService struct {
	SpuRepo repo.ISpuRepo
}

func NewSpuService(spuRepo repo.ISpuRepo) *SpuService {
	svc := SpuService{SpuRepo: spuRepo}
	return &svc
}

//  CreateSpu create new spu
func (s *SpuService) CreateSpu(req *entity.CreateSpuReq) (res *entity.CreateSpuReq, err error) {
	return res, nil
}

// getSpuInfo 获取商品的信息
func (s *SpuService) GetSpuInfo(req *entity.GetSpuInfoReq) (res *entity.GetSpuInfoRes, err error) {
	// 直接查数据然后返回
	spuInfo, err := s.SpuRepo.GetSpuDetail(context.TODO(), req.AppID, req.SpuID)
	if err != nil {
		return nil, quickErrors.New(consts.CurdSelectErrorCode, consts.CurdSelectErrorMsg, err.Error())
	}

	if spuInfo.AppID == "" {
		return nil, quickErrors.New(consts.CurdSelectErrorCode, consts.CurdSelectErrorMsg, "spu not found")
	}

	res = &entity.GetSpuInfoRes{
		AppID:        spuInfo.AppID,
		SpuID:        spuInfo.SpuID,
		SpuType:      spuInfo.SpuType,
		ResourceID:   spuInfo.ResourceID,
		ResourceType: spuInfo.ResourceType,
		GoodsName:    spuInfo.GoodsName,
		GoodsImg:     spuInfo.GoodsImg,
	}

	resDataStr, _ := json.Marshal(res)
	global.KafkaProLocal.Input() <- &sarama.ProducerMessage{
		Topic: "revolution",
		Key:   sarama.ByteEncoder("spu_datas"),
		Value: sarama.ByteEncoder(resDataStr),
	}

	return res, nil
}

// getSpuSaleTimeStamp 获取商品的销售时间戳
func getSpuSaleTimeStamp(saleAtString string) (saleAtStamp int64, err error) {
	saleAt, err := time.Parse("2006-01-02 15:04:05", saleAtString)
	if err != nil {
		return 0, err
	}
	return saleAt.Unix(), nil
}

func (s *SpuService) AsyncRedisList(req *entity.GetSpuInfoReq) (res *entity.GetSpuInfoRes, err error) {
	// 直接拿参数出来写到redis 队列上面
	reqStr, _ := json.Marshal(req)

	_, err = global.RedisLocal.LPush("goods_detail_list", reqStr).Result()
	if err != nil {
		fmt.Println(err.Error())
	}

	return res, nil
}
