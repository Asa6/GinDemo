package middleware

import (
	. "GinDemo/common"
	. "GinDemo/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			JsonResponse(c, http.StatusBadRequest, INVALID_PARAMS, GetMsg(INVALID_PARAMS), nil)
		} else {
			val, err := RDB.Get(Ctx, "token_"+token).Result()
			if err != nil {
				JsonResponse(c, http.StatusBadRequest, ERROR_AUTH_CHECK_TOKEN_FAIL, GetMsg(ERROR_AUTH_CHECK_TOKEN_FAIL), nil)
			} else {
				// 向业务接口传递通过token认证的登录用户名
				c.Set("Login_UserName", val)

				// before request
				c.Next()
				// after request
			}
		}
	}
}
