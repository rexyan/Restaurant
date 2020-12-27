package dao

import (
	"Restaurant/model"
	"Restaurant/tool"
	"log"
)

type MemberDao struct {
	*tool.Orm
}

// AddMember 新增 Member
func (md *MemberDao) AddMember(member *model.Member) *model.Member {
	id, err := md.InsertOne(member)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	member.Id = id
	return member
}

// GetMemberByMobile 根据用户 mobile 查询用户信息
func (md *MemberDao) GetMemberByMobile(mobile string) *model.Member {
	var result model.Member
	_, err := md.Where("mobile = ?", mobile).Get(&result)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &result
}
