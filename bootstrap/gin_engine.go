package bootstrap

import "github.com/gin-gonic/gin"

func NewGinEngine() *gin.Engine {
	return gin.Default()
}
