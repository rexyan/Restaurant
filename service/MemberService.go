package service

import (
	"Restaurant/tool"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type MemberService struct {
}

func (ms *MemberService) SendCode(phone string) bool {
	smsConfig := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(smsConfig.RegionId, smsConfig.AppKey, smsConfig.AppSecret)
	if err!= nil{

		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = phone
	request.SignName = smsConfig.SignName
	request.TemplateCode = smsConfig.TemplateCode
	request.Scheme = smsConfig.Schema

	if response, err := client.SendSms(request); err != nil {
		return false
	} else {
		if response.Code == "OK" {
			return true
		}
		return false
	}
}
