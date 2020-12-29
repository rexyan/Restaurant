package service

import (
	"Restaurant/dao"
	"Restaurant/model"
)

type FoodCategoryService struct {
}

func (fc *FoodCategoryService) GetFoodCategory() []model.FoodCategory {
	foodCategoryDao := dao.FoodCategoryDao{}
	return foodCategoryDao.GetFoodCategory()
}
