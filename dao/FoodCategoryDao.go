package dao

import (
	"Restaurant/model"
	"Restaurant/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

func (fc *FoodCategoryDao) GetFoodCategory() []model.FoodCategory {
	var categories []model.FoodCategory
	if err := fc.Find(&categories); err != nil {
		return nil
	}
	return categories
}
