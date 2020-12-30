package model

//服务结构体定义
type Service struct {
	//id
	Id int64 `xorm:"pk autoincr" json:"id"`
	//服务名称
	Name string `xorm:"varchar(20)" json:"name"`
	//服务描述
	Description string `xorm:"varchar(30)" json:"description"`
	//图标名称
	IconName string `xorm:"varchar(3)" json:"icon_name"`
	//图标颜色
	IconColor string `xorm:"varchar(6)" json:"icon_color"`
}
