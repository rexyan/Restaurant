package service

import (
	"Restaurant/tool"
	"fmt"
	"math/rand"
	"time"
)

type MemberService struct {
}

// SendCode 发送验证码
func (ms *MemberService) SendCode(phone string) bool {
	// 生成短信随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 发送信息
	esandSms := tool.EsandSms{}
	redisStore := tool.RedisStoreEngine
	redisStore.Set(phone, code, time.Minute*5)
	return esandSms.SendSms(phone, code)
}

// CheckSmsCode 校验验证码
func (ms *MemberService) CheckSmsCode(phone string, code string) bool {
	redisStore := tool.RedisStoreEngine
	return redisStore.Get(phone) == code
}
