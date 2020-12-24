package service

import (
	"Restaurant/dao"
	"Restaurant/model"
	"Restaurant/tool"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) SendCode(phone string) bool {
	// 生成短信随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 获取 sms 配置
	smsConfig := tool.GetConfig().Sms
	// 验证阿里云 ak 信息
	client, err := dysmsapi.NewClientWithAccessKey(smsConfig.RegionId, smsConfig.AppKey, smsConfig.AppSecret)
	if err != nil {
		return false
	}
	// 创建发送 sms request
	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = phone
	request.SignName = smsConfig.SignName
	request.TemplateCode = smsConfig.TemplateCode
	request.Scheme = smsConfig.Schema
	templateParam, _ := json.Marshal(gin.H{
		"code": code,
	})
	request.TemplateParam = string(templateParam)
	// 发送信息
	if response, err := client.SendSms(request); err != nil {
		return false
	} else {
		if response.Code == "OK" {
			// code 记录在数据库中
			smsCode := model.SmsCode{
				Phone:      phone,
				BizId:      response.BizId,
				Code:       code,
				CreateTime: time.Now().Unix(),
			}
			memberDao := dao.MemberDao{tool.DBEngine}
			return memberDao.AddSmsCode(smsCode) > 0
		}
		return false
	}
}
