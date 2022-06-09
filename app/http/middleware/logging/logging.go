package logging

import (
	"bytes"
	"io/ioutil"
	"quick-go/global"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/// LoggingWriter ...
type LoggingWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &LoggingWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}

		logFields := []zap.Field{}

		// 抽取参数
		c.Writer = w // hijack 走一波
		requestBody, _ := ioutil.ReadAll(c.Request.Body)
		// 填回参数至 c.Request.Body, 因为上一步的抽取参数会将其覆写为空
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		// 记录 Log 内容
		st := time.Now()
		logFields = append(
			logFields,
			zap.ByteString("req", requestBody),
			zap.String("client_ip", strings.Split(c.Request.RemoteAddr, ":")[0]),
			zap.String("target_url", c.Request.RequestURI),
		)

		// 请求后
		c.Next()

		// 获取参数
		responseBody := strings.Trim(w.body.String(), "\n")
		// 记录 Log 内容
		logFields = append(
			logFields,
			zap.Int("http_code", w.Status()),
			zap.String("ret", responseBody),
			zap.String("ts_end", time.Now().Format(time.RFC3339)),
			zap.Float64("latency", time.Since(st).Seconds()),
		)

		global.RequestLogger.Info("Middleware logging", logFields...)

	}
}
