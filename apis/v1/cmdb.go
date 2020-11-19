package v1

import (
	. "GinDemo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCmdb(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  GetMsg(SUCCESS),
		"data": nil,
	})
}

func AddCmdb(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": "",
		"msg":  "",
		"data": nil,
	})
}
