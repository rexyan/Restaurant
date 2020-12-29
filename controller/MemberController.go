package controller

import (
	"Restaurant/dao"
	"Restaurant/enums"
	"Restaurant/model"
	"Restaurant/service"
	"Restaurant/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MemberController struct {
	BaseController
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendSms", mc.sendSmsCode)
	engine.POST("/api/smsLogin", mc.smsLogin)
	engine.GET("/api/captcha", mc.captcha)
	engine.POST("/api/login", mc.login)
}

//sendSmsCode 发送验证码
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist {
		BuildResponse(context, http.StatusBadRequest, enums.SMSError, "params 'phone' not exist")
		return
	}
	memberService := service.MemberService{}
	sendStatus := memberService.SendCode(phone)
	if sendStatus {
		BuildSuccessResponse(context, "send sms success!")
		return
	}
	BuildResponse(context, http.StatusBadRequest, enums.SMSError, "send sms error!")
	return
}

// smsLogin 短信登陆，只要验证码验证成功就登陆，没账号就自动注册
func (mc *MemberController) smsLogin(context *gin.Context) {
	requestData := RequestData(context)
	phone := requestData["phone"]
	code := requestData["code"]
	if phone == "" || code == "" {
		BuildResponse(context, http.StatusBadRequest, enums.ParamError, "phone and code is a required parameter")
		return
	}
	memberService := service.MemberService{}
	if !memberService.CheckSmsCode(phone, code) {
		BuildResponse(context, http.StatusBadRequest, enums.ParamError, "sms code error!")
		return
	}
	memberDao := dao.MemberDao{Orm: tool.DBEngine}
	member := memberDao.GetMemberByMobile(phone)
	if member.Id == 0 {
		member = &model.Member{
			UserName:     phone,
			Mobile:       phone,
			Password:     tool.Md5("123456"),
			RegisterTime: time.Now().Unix(),
			IsActive:     1,
		}
		member = memberDao.AddMember(member)
	}
	sessionValue := map[string]interface{}{"Mobile": phone, "Avatar": member.Avatar, "Balance": member.Balance, "City": member.City}
	tool.SetSession(phone, sessionValue)
	BuildResponse(context, http.StatusOK, enums.SUCCESS, member)
}

// captcha 获取验证码
func (mc *MemberController) captcha(context *gin.Context) {
	id, b64s := tool.GenerateCaptcha()
	BuildResponse(context, http.StatusOK, enums.SUCCESS, map[string]string{"id": id, "b64s": b64s})
}

// login 账号密码登陆
func (mc *MemberController) login(context *gin.Context) {
	requestData := RequestData(context)
	phone := requestData["name"]
	pwd := requestData["pwd"]
	captchaId := requestData["id"]
	captchaValue := requestData["value"]
	if !tool.CheckCaptcha(captchaId, captchaValue) {
		BuildResponse(context, http.StatusBadRequest, enums.CaptchaError, "captcha code error!")
		return
	}
	memberService := service.MemberService{}
	member := memberService.GetMemberByMobile(phone)
	if !memberService.MemberExist(member) || !memberService.CheckMemberPassword(member, pwd) {
		BuildResponse(context, http.StatusBadRequest, enums.UnregisteredOrPasswordError, "phone number unregistered or password error!")
		return
	}
	sessionValue := map[string]interface{}{"Mobile": phone, "Avatar": member.Avatar, "Balance": member.Balance, "City": member.City}
	tool.SetSession(phone, sessionValue)
	BuildResponse(context, http.StatusOK, enums.SUCCESS, member)
}
