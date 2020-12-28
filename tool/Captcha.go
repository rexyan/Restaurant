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
	driver := base64Captcha.NewDriverDigit(50, 120, 4, 0.7, 50)
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		id = ""
		b64s = ""
	}
	return id, b64s
}

func CheckCaptcha(id string, value string) bool {
	if store.Verify(id, value, true) {
		return true
	}
	return false
}
