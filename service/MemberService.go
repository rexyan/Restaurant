package service

import (
	"Restaurant/dao"
	"Restaurant/model"
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

// CheckSmsCode 校验短信验证码
func (ms *MemberService) CheckSmsCode(phone string, code string) bool {
	redisStore := tool.RedisStoreEngine
	return redisStore.Get(phone) == code
}

// GetMemberByMobile 根据 Mobile 查询 member 信息
func  (ms *MemberService) GetMemberByMobile(phone string) *model.Member {
	memberDao := dao.MemberDao{Orm: tool.DBEngine}
	return memberDao.GetMemberByMobile(phone)
}

// MemberExist 根据传入 Member 对象，判断 Member 是否合法
func (ms *MemberService) MemberExist(member *model.Member) bool {
	return member.Id >0
}

func (ms *MemberService) CheckMemberPassword(member *model.Member, plaintext string) bool {
	return member.Password == tool.Md5(plaintext)
}
