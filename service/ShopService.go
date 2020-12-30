package service

import (
	"Restaurant/dao"
	"Restaurant/model"
)

type ShopService struct {
}

func (sp *ShopService) GetShopByPosition(longitude, latitude float64) []model.Shop {
	shopDao := dao.ShopDao{}
	return shopDao.GetShopByPosition(longitude, latitude)
}

func (sp *ShopService) SearchShop(longitude, latitude float64, keyword string) []model.Shop {
	shopDao := dao.ShopDao{}
	return shopDao.SearchShop(longitude, latitude, keyword)
}