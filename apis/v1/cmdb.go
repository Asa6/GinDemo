package v1

import (
	"github.com/gin-gonic/gin"
)

func GET_CMDB_API(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "cmdb api",
	})
}
