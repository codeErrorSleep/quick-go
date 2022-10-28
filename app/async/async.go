package async

import (
	"fmt"
	"quick-go/global"
	"time"

	"github.com/go-redis/redis"
)

func AsyncGoodsDetail() {
	for {
		ret, err := global.RedisLocal.BLPop(5*time.Second, "goods_detail_list").Result()
		fmt.Print(ret)
		if err != nil {
			if err == redis.Nil {
				fmt.Print("正常")
			} else {
				fmt.Print(err.Error())
			}
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Print(ret)
		time.Sleep(50 * time.Millisecond)
	}
}
