package controller

import (
	"Restaurant/service"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {
	BaseController
}

func (fc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/foodCategory", fc.GetFoodCategory)
}

func (fc *FoodCategoryController) GetFoodCategory(context *gin.Context) {
	foodCategoryService := service.FoodCategoryService{}
	foodCategory := foodCategoryService.GetFoodCategory()
	BuildSuccessResponse(context, foodCategory)
}
