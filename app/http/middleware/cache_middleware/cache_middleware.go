package cache_middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quick-go/app/entity"
	"quick-go/global"
	"quick-go/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type cacheWrappedWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *cacheWrappedWriter) Write(body []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(body)
	if err == nil {
		rw.body.Write(body)
	}
	return n, err
}

func SetRedisCache(expireTime time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 抽取参数
		// 默认http.Request.Body类型为io.ReadCloser类型,即只能读一次，读完后直接close掉
		// 填回参数至 c.Request.Body, 因为上一步的抽取参数会将其覆写为空
		requestBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		var requestUrl string
		endIndex := strings.Index(c.Request.RequestURI, "?")
		if endIndex > 0 {
			requestUrl = c.Request.RequestURI[1:endIndex]
		} else {
			requestUrl = c.Request.RequestURI[1:]
		}
		fmt.Print(expireTime)
		// 查redis 缓存
		cacheKeyMd5 := requestUrl + string(requestBody)
		cachekey := "quick_middleware_" + utils.LiteralToMD5(cacheKeyMd5)
		fmt.Print(cachekey)
		cacheValue := global.RedisLocal.Get(cachekey)
		if cacheValue.Val() != "" {
			ret := entity.DefaultRequest{}
			_ = json.Unmarshal([]byte(cacheValue.Val()), &ret)
			c.AbortWithStatusJSON(http.StatusOK, ret)
			return
		}

		temp := c.Writer
		w := &cacheWrappedWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		// 查到返回,
		c.Next()
		c.Writer = temp

		// 只缓存成功的
		if !c.IsAborted() && w.Status() < 300 && w.Status() >= 200 {
			if err := global.RedisLocal.Set(cachekey, w.body.String(), time.Second*expireTime); err != nil {
				global.ErrorLogger.Info("cache middleware failed	", zap.Error(err.Err()))
			}
		}
	}
}
