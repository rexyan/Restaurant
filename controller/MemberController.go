package controller

import (
	"Restaurant/enums"
	"Restaurant/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendSms", mc.sendSmsCode)
}

//sendSmsCode 发送验证码
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist {
		BuildResponse(context, http.StatusBadRequest,enums.SMSError, "params 'phone' not exist")
		return
	}
	memberService := service.MemberService{}
	sendStatus := memberService.SendCode(phone)
	if sendStatus{
		BuildSuccessResponse(context, "send sms success!")
		return
	}
	BuildResponse(context, http.StatusBadRequest, enums.SMSError, "send sms error!")
	return
}

// smsLogin 短信登陆，只要验证码验证成功就登陆，没账号就自动注册
func (mc *MemberController) smsLogin(context *gin.Context)  {
	phone := context.PostForm("phone")
	code := context.PostForm("code")
	if phone == "" || code == ""{
		BuildResponse(context, http.StatusBadRequest, enums.ParamError, "phone and code is a required parameter")
		return
	}
	memberService := service.MemberService{}
	if !memberService.CheckSmsCode(phone, code){
		BuildResponse(context, http.StatusBadRequest, enums.ParamError, "phone and code is a required parameter")
		return
	}
}
