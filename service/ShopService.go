package service

import (
	"Restaurant/dao"
	"Restaurant/model"
	"Restaurant/tool"
)

type ShopService struct {
}

// GetShopByPosition
func (sp *ShopService) GetShopByPosition(longitude, latitude float64) []model.Shop {
	shopDao := dao.ShopDao{Orm: tool.DBEngine}
	return shopDao.GetShopByPosition(longitude, latitude)
}

// SearchShop
func (sp *ShopService) SearchShop(longitude, latitude float64, keyword string) []model.Shop {
	shopDao := dao.ShopDao{Orm: tool.DBEngine}
	return shopDao.SearchShop(longitude, latitude, keyword)
}

// GetShopService
func (sp *ShopService) GetShopService(shopId int64) [] model.Service{
	shopDao := dao.ShopDao{Orm: tool.DBEngine}
	return shopDao.GetShopService(shopId)
}