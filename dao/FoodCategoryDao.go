package dao

import (
	"Restaurant/model"
	"Restaurant/tool"
	"fmt"
)

type FoodCategoryDao struct {
	*tool.Orm
}

func (fc *FoodCategoryDao) GetFoodCategory() []model.FoodCategory {
	var categories []model.FoodCategory
	if err := fc.Engine.Find(&categories); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return categories
}
