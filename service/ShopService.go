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
