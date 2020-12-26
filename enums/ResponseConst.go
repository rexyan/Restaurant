package enums

const (
	SUCCESS    string = "90000"
	ParamError string = "30011"
	SMSError   string = "30012"
)

var ErrorMessage = map[string]string{
	SUCCESS:    "成功",
	ParamError: "参数错误",
	SMSError:   "SMS 错误",
}
