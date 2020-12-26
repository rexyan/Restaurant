package dao

import (
	"Restaurant/model"
	"Restaurant/tool"
	"log"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) AddSmsCode(member model.Member) int64{
	rows, err := md.InsertOne(&member)
	if err!= nil{
		log.Println(err.Error())
	}
	return rows
}