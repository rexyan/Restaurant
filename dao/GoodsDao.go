package dao

import (
	"Restaurant/model"
	"Restaurant/tool"
)

type GoodsDao struct {
	*tool.Orm
}

// GetGoodsByShopId
func (gd *GoodsDao) GetGoodsByShopId(shopId int64) []model.Goods {
	var goods []model.Goods
	_, err := gd.Where("shop_id = ?", shopId).Get(&goods)
	if err != nil {
		return nil
	}
	return goods
}
