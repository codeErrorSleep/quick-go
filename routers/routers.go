package routers

import (
	"net/http"
	"quick-go/app/http/controllers"
	cors2 "quick-go/app/http/middleware/cors"
	logging2 "quick-go/app/http/middleware/logging"

	"github.com/gin-gonic/gin"
)

type Router struct {
	stock controllers.StockController
}

func NewRouter(stockController controllers.StockController) *Router {
	router := &Router{
		stock: stockController,
	}
	return router
}

func (r *Router) With(engine *gin.Engine) {
	engine.POST("/reduce_stock", r.stock.ReduceStock)
	engine.POST("/get_spu_stock", r.stock.GetSpuStock)
}

// InitApiRouter ...
func InitApiRouter(test bool) *gin.Engine {
	router := gin.Default()

	//全局中间件
	if test {
		gin.SetMode("test")
	} else {
		router.Use(logging2.Logging()) //日志中间件
		router.Use(cors2.CORS())       //跨域中间件
	}

	//探针
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello logistics")
	})

	// stockAPI := router.Group("/stock/")
	// {
	// 	stockAPI.POST("/reduce_stock", controllers.ReduceStock)
	// 	stockAPI.POST("/get_spu_stock", controllers.GetSpuStock)
	// 	// stockAPI.POST("/get_spu_stock", cache_middleware.SetRedisCache(100), controllers.GetSpuStock)
	// }

	// spuAPI := router.Group("/spu/")
	// {
	// 	spuAPI.POST("/get_spu_info", controllers.GetSpuInfo)
	// 	spuAPI.POST("/create_test_val", controllers.CreateSpu)
	// }

	// kafkaAPI := router.Group("/kafka/")
	// {
	// 	kafkaAPI.POST("createdManyMessages", controllers.CreateManyMessages)
	// }

	// AsyncAPI := router.Group("/async/")
	// {
	// 	AsyncAPI.POST("async_redis_list", controllers.AsyncRedisList)
	// }

	return router
}
