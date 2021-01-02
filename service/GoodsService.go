package service

import (
	"Restaurant/dao"
	"Restaurant/model"
	"Restaurant/tool"
)

type GoodsService struct {
}

// GetGoodsByShopId
func (gs *GoodsService) GetGoodsByShopId(shopId int64) []model.Goods {
	goodsDao := dao.GoodsDao{Orm: tool.DBEngine}
	return goodsDao.GetGoodsByShopId(shopId)
}
