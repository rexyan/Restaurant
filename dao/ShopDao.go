package dao

import (
	"Restaurant/model"
	"Restaurant/tool"
)

type ShopDao struct {
	*tool.Orm
}

// GetShopByPosition, 根据经纬度查询 Shop 信息
func (sp *ShopDao) GetShopByPosition(longitude, latitude float64) []model.Shop {
	var shops []model.Shop
	_, err := sp.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and status = 1", longitude, latitude).Get(&shops)
	if err != nil {
		return nil
	}
	return shops
}

// SearchShop 根据经纬度，关键字查询商家
func (sp *ShopDao) SearchShop(longitude, latitude float64, keyword string) []model.Shop {
	var shops []model.Shop
	_, err := sp.Where("longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status = 1", longitude, latitude, keyword).Get(&shops)
	if err != nil {
		return nil
	}
	return shops
}

// GetShopService 查询商家拥有的服务
func (sp *ShopDao) GetShopService(shopId int64) []model.Service {
	var services []model.Service
	err := sp.Table("service").Join("inner", "shop_service", "service.id=shop_service.service_id and shop_service.ship=?", shopId).Find(&services)
	if err != nil {
		return nil
	}
	return services
}
