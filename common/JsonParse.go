package common

import (
	. "GinDemo/model"
	"github.com/gin-gonic/gin"
)

// 统一的json解析组件
func JsonParse(c *gin.Context, u *User) *User {
	c.BindJSON(&u)
	return u
}
