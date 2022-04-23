package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	r *gin.Engine
)


func TestGetSpuStock(t *testing.T){
	// 发起请求
	w := request(r, "GET", "/", nil)

	// 检查结果
	assert.Equal(t, 200, w.Code)
}



