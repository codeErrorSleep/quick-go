package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"quick-go/app/entity"
	"quick-go/global"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
)

type KafkaService struct {
	ctx *gin.Context
}

func KafaServiceNew(ctx *gin.Context) *KafkaService {
	return &KafkaService{
		ctx: ctx,
	}
}

// CreateManyMessages 创建多条kafka的message
func (ks *KafkaService) CreateManyMessages() (err error) {
	for i := 0; i < 100; i++ {
		s := fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(100000000))

		spuInfo := entity.GetSpuInfoRes{
			AppID: s,
			SpuID: s,
		}
		spuInfoStr, _ := json.Marshal(spuInfo)
		global.KafkaProLocal.Input() <- &sarama.ProducerMessage{
			Topic: "revolution22",
			Key:   sarama.ByteEncoder("spu_datas"),
			Value: sarama.ByteEncoder(spuInfoStr),
		}
	}
	return nil
}
