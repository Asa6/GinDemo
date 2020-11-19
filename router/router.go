package router

import (
	. "GinDemo/apis/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/cmdbs", GET_CMDB_API)
	return r
}
