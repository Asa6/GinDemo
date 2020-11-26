package v1

import (
	. "GinDemo/common"
	. "GinDemo/config"
	. "GinDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	// 获取此登录用户名
	login_user, err := c.MustGet("Login_UserName").(string)
	if err {
		// 根据Authorization中间件返回的用户，查询用户具体信息
		var u User
		DB.Where(&User{UserName: login_user}).First(&u)
		data := map[string]interface{}{"username": u.UserName, "id": u.ID, "email": u.Email, "fullname": u.FullName}

		JsonResponse(c, http.StatusOK, SUCCESS, GetMsg(SUCCESS), data)
	} else {
		JsonResponse(c, http.StatusBadRequest, INVALID_PARAMS, GetMsg(INVALID_PARAMS), nil)
	}
}
