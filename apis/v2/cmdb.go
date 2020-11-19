package v2

import (
	"github.com/gin-gonic/gin"
)

func GetCmdb(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "cmdb v2 api",
	})
}
