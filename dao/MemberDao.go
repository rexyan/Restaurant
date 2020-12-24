package dao

import (
	"Restaurant/model"
	"Restaurant/tool"
	"fmt"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) AddSmsCode(sms model.SmsCode) int64{
	rows, err := md.InsertOne(&sms)
	if err!= nil{
		fmt.Println(err.Error())
	}
	return rows
}