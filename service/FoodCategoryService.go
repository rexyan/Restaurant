package service

import (
	"Restaurant/dao"
	"Restaurant/model"
	"Restaurant/tool"
)

type FoodCategoryService struct {
}

func (fc *FoodCategoryService) GetFoodCategory() []model.FoodCategory {
	foodCategoryDao := dao.FoodCategoryDao{Orm: tool.DBEngine}
	return foodCategoryDao.GetFoodCategory()
}
