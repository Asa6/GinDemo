package main

import (
	. "GinDemo/config"
	. "GinDemo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 设置gin的模式
	gin.SetMode(gin.DebugMode)

	// 获取数据库环境信息并连接
	var d Database
	dbInfo := d.GetInfo()
	db := dbInfo.GetConnect()

	// 检测表模型是否被创建
	CreateModel(db)

	// 关闭链接
	defer db.Close()

	// 路由挂载
	r := InitRouter()
	r.Run(":8999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
