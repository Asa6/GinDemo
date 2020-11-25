package common

import (
	"github.com/gin-gonic/gin"
)

// 统一的json返回组件
func JsonResponse(c *gin.Context, HTTPCode int, code int, msg string) {
	c.JSON(HTTPCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}
