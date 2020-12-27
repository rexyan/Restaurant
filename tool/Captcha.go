package tool

import (
	"github.com/mojocn/base64Captcha"
)

type ConfigJsonBody struct {
	Id          string
	CaptchaType string
	VerifyValue string
	Driver      *base64Captcha.DriverMath
}

var store = base64Captcha.DefaultMemStore

func GenerateCaptcha() (string, string) {
	driver := base64Captcha.DefaultDriverDigit
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		id = ""
		b64s = ""
	}
	return id, b64s
}

func CheckCaptcha(configJsonBody ConfigJsonBody) bool {
	if store.Verify(configJsonBody.Id, configJsonBody.VerifyValue, true) {
		return true
	}
	return false
}
