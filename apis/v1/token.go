package v1

import (
	. "GinDemo/common"
	. "GinDemo/config"
	. "GinDemo/model"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var ctx = context.Background()

func AddToken(c *gin.Context) {
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
		JsonResponse(c, http.StatusUnauthorized, INVALID_PARAMS, GetMsg(INVALID_PARAMS), nil)
	} else {
		// 根据登录用户名和当前的时间生成md5 token
		m5 := md5.New()
		m5.Write([]byte(u.UserName))
		m5.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
		token := hex.EncodeToString(m5.Sum(nil))

		// 缓存进redis，7天自动过期
		err := RDB.Set(ctx, "token_"+token, u.UserName, 7*24*time.Hour).Err()
		if err != nil {
			//panic(err)
			fmt.Println("token申请失败：", err)
		}

		//val, err := RDB.Get(ctx, "key").Result()
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println("key", val)

		JsonResponse(c, http.StatusOK, SUCCESS, GetMsg(SUCCESS), map[string]string{"token": token})
	}
}
