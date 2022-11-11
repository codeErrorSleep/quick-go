package test

import (
	"bytes"
	"encoding/json"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"quick-go/bootstrap"
	"quick-go/global/consts"
	"quick-go/routers"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	r           *gin.Engine
	initialized bool
)

func TestMain(m *testing.M) {
	flag.Parse() // 读取命令行参数
	os.Exit(m.Run())
}

func initTest() {
	// 初始化工程
	bootstrap.Bootstrap(consts.EnvUnitTest)
	initialized = true
	r = routers.InitApiRouter(true)
	// go r.Run(":" + global.Env.GetString("httpPort"))
}

func TestAll(t *testing.T) {
	// 初始化工程
	initTest()

	time.Sleep(time.Second)
	t.Run("Default", func(t *testing.T) {
		t.Run("healthCheck", testHealthCheck)
	})

}

// testHealthCheck 测试服务是否启动正常
func testHealthCheck(t *testing.T) {
	ret := request(r, "GET", "/", nil)
	assert.Equal(t, http.StatusOK, ret.Code)
}

// request sends out a request and returns the response
func request(r http.Handler, method, path string, reqData interface{}) *httptest.ResponseRecorder {
	argsBytes, _ := json.Marshal(reqData)
	req, _ := http.NewRequest(method, path, bytes.NewReader(argsBytes))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
