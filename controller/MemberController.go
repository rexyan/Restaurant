package controller

import (
	"Restaurant/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {

}

func (mc *MemberController) Router(engine *gin.Engine)  {
	engine.GET("/api/sendcode", mc.sendSmsCode)
}

func (mc *MemberController) sendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist{
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"message": "参数解析失败",
		})
		return
	}
	ms := service.MemberService{}
	isSend := ms.SendCode(phone)
	if isSend{
		context.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "发送成功",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "发送失败",
	})
}