package middleware

import "github.com/gin-gonic/gin"

func Auth(c *gin.Context) {
	token := "123"
	if token == "123" {
		panic("token not exist !")
	}
}
