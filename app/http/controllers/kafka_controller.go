package controllers

import (
	"fmt"
	"quick-go/app/service"

	"github.com/gin-gonic/gin"
)

// CreateManyMessages 生成kafka多个massage
func CreateManyMessages(c *gin.Context) {
	svc := service.KafaServiceNew(c)
	err := svc.CreateManyMessages()
	fmt.Print(err)
	// response.Respond(c, data, err)
}
