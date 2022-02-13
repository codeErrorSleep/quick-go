package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quick-go/app/http/controllers"
	cors2 "quick-go/app/http/middleware/cors"
	logging2 "quick-go/app/http/middleware/logging"
)

// InitApiRouter ...
func InitApiRouter(test bool) *gin.Engine {
	router := gin.Default()

	//全局中间件
	if !test {
		router.Use(logging2.Logging()) //日志中间件
		router.Use(cors2.CORS())       //跨域中间件
	}

	//探针
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello logistics")
	})

	stockAPI := router.Group("/stock/")
	{
		stockAPI.POST("/reduce_stock", controllers.ReduceStock)
		stockAPI.POST("/get_spu_stock", controllers.GetSpuStock)
	}

	spuAPI := router.Group("/spu/")
	{
		spuAPI.POST("/get_spu_info", controllers.GetSpuInfo)

	}

	return router
}
