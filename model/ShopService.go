package model

type ShopService struct {
	ShopId    int64 `xorm:"pk not null" json:"shop_id"`
	ServiceId int64 `xorm:"pk not null" json:"service_id"`
}
