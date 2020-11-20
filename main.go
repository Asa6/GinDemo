package main

import (
	. "GinDemo/config"
	. "GinDemo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 设置gind模式
	gin.SetMode(gin.DebugMode)

	var d Database
	dbInfo := d.GetInfo()
	dbInfo.GetConnect()
	println(dbInfo.Password)
	println(dbInfo.Username)
	println(dbInfo.Database)

	// 路由挂载
	r := InitRouter()
	r.Run(":8999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
