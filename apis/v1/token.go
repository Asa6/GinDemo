package v1

import (
	. "GinDemo/common"
	. "GinDemo/config"
	. "GinDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetToken(c *gin.Context) {
	// 解析json请求
	u := JsonParse(c, &User{})

	// 获取用户名和密码
	username := u.UserName
	password := u.PassWord

	// 查询用户和密码是否存在
	var user User
	DB.Where(&User{UserName: username, PassWord: password}).First(&user)

	// 判断查询结果是否为空
	if user == (User{}) {
		JsonResponse(c, http.StatusUnauthorized, INVALID_PARAMS, GetMsg(INVALID_PARAMS))
	} else {
		JsonResponse(c, http.StatusOK, SUCCESS, GetMsg(SUCCESS))
	}
}
