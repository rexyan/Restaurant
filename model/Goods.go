package model

// 食品
type Goods struct {
	Id          int64   `xorm:"pk autoincr" json:"id"`          //商品Id
	Name        string  `xorm:"varchar(12)" json:"name"`        //商品名称
	Description string  `xorm:"varchar(32)" json:"description"` //商品描述
	Icon        string  `xorm:"varchar(255)" json:"icon"`       //商品图标
	SellCount   int64   `xorm:"int" json:"sell_count"`          //销售份数
	Price       float32 `xorm:"float" json:"price"`             //销售价格
	OldPrice    float32 `xorm:"float" json:"old_price"`         //原价
	ShopId      int64   `xorm:"int" json:"shop_id"`             //商品ID，表明该商品属于哪个商家
}
