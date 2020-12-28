package enums

const (
	SUCCESS       string = "90000"
	ParamError    string = "30011"
	SMSError      string = "30012"
	CaptchaError  string = "30013"
	UnregisteredOrPasswordError  string = "30014"
	Unauthorized  string = "30015"
)

var ErrorMessage = map[string]string{
	SUCCESS:       "成功",
	ParamError:    "参数错误",
	SMSError:      "SMS 错误",
	CaptchaError:  "验证码错误",
	UnregisteredOrPasswordError: "未注册或者密码错误",
	Unauthorized: "未登录",
}
