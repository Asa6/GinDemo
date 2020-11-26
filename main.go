package main

import (
	. "GinDemo/config"
	. "GinDemo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode) // 设置gin的模式

	var d DBEnv
	dbInfo := d.GetInfo() // 获取数据库环境信息并连接
	dbInfo.GetConnect()   // 连接数据库

	//CreateModel()				// 检测表模型是否被创建
	defer DB.Close() // 关闭数据库链接

	var rn RedisENV
	redisInfo := rn.GetInfo()  // 获取redis环境信息
	redisInfo.GetRedisClient() // 连接redis

	// 路由挂载
	r := InitRouter()
	r.Run(":8999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
