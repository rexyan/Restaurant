package middleware

import (
	"Restaurant/controller"
	"Restaurant/enums"
	"Restaurant/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginRequiredMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 从三个地方获取 phone 的值
		var phone string
		phone = context.DefaultPostForm("phone", "")
		if phone == ""{
			phone = context.DefaultQuery("phone", "")
		}
		if phone == ""{
			phone, _ = context.Cookie("phone")
		}
		// 查询 redis 看当前 phone 是否登录
		userInfo := tool.GetSession(phone)
		if len(userInfo) == 0{
			controller.BuildResponse(context, http.StatusUnauthorized, enums.Unauthorized, "")
			return
		}
		context.Set("userInfo", userInfo)
	}
}