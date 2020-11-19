package main

//import "github.com/gin-gonic/gin"

import . "GinDemo/router"

func main() {
	// 路由挂载
	r := InitRouter()
	r.Run(":8999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
