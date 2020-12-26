package tool

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type EsandSms struct {
}

func (es EsandSms) SendSms(phone string, code string) bool {
	// 获取 sms 配置
	smsConfig := GetConfig().Sms

	urlValues := url.Values{}
	urlValues.Add("mobile", "+86"+phone)
	urlValues.Add("templateID", smsConfig.TemplateID)
	urlValues.Add("templateParamSet", code+", 1")

	body := strings.NewReader(urlValues.Encode())
	req, _ := http.NewRequest("POST", smsConfig.URL, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "APPCODE "+smsConfig.AppCode)
	clt := http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	var respBody map[string]interface{}
	json.Unmarshal(content, &respBody)
	if respBody["code"] == "0000" {
		return true
	}
	return false
}
