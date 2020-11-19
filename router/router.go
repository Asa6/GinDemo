package router

import (
	v1 "GinDemo/apis/v1"
	v2 "GinDemo/apis/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 使用gin Group进行路由分组
	ApiV1 := r.Group("/api/v1")
	{
		ApiV1.GET("/cmdbs", v1.GET_CMDB_API)
	}

	ApiV2 := r.Group("/api/v2")
	{
		ApiV2.GET("/cmdbs", v2.GET_CMDB_API)
	}

	return r
}
