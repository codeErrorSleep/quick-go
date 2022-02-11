package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quick-go/http/middleware/cors"
	"quick-go/http/middleware/logging"
)

// InitApiRouter ...
func InitApiRouter(test bool) *gin.Engine {
	router := gin.Default()

	//全局中间件
	if !test {
		router.Use(logging.Logging()) //日志中间件
		router.Use(cors.CORS())       //跨域中间件
	}

	//探针
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello logistics")
	})

	return router
}
