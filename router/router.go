package router

import (
	v1 "GinDemo/apis/v1"
	v2 "GinDemo/apis/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(Auth)

	// 使用gin Group进行路由分组
	ApiV1 := r.Group("/api/v1")
	{
		// cmdb相关接口
		ApiV1.GET("/cmdbs/", v1.GetCmdb)
		ApiV1.POST("/cmdbs/", v1.AddCmdb)

		// 登录认证相关接口
		ApiV1.POST("/token/", v1.AddToken)
	}

	ApiV2 := r.Group("/api/v2")
	{
		ApiV2.GET("/cmdbs/", v2.GetCmdb)
	}

	return r
}
