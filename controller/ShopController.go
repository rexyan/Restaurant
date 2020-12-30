package controller

import (
	"Restaurant/enums"
	"Restaurant/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ShopController struct {
	BaseController
}

func (sp *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/getShopByPosition", sp.GetShopByPosition)
	engine.GET("/api/searchShop", sp.SearchShop)
}

// GetShopByPosition 根据经纬度查询商家列表
func (sp *ShopController) GetShopByPosition(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	longitudeFloat, err1 := strconv.ParseFloat(longitude, 10)
	latitudeFloat, err2 := strconv.ParseFloat(latitude, 10)
	if err1 != nil || err2 != nil {
		BuildResponse(context, http.StatusBadRequest, enums.PositionError, "params 'longitude' or 'latitude' error!")
		return
	}
	shopService := service.ShopService{}
	shopByPosition := shopService.GetShopByPosition(longitudeFloat, latitudeFloat)
	BuildSuccessResponse(context, shopByPosition)
}

// SearchShop 根据经纬度, 关键词搜索商家
func (sp *ShopController) SearchShop(context *gin.Context) {
	longitude := context.Query("longitude")
	latitude := context.Query("latitude")
	keyword := context.DefaultQuery("keyword", "")
	longitudeFloat, err1 := strconv.ParseFloat(longitude, 10)
	latitudeFloat, err2 := strconv.ParseFloat(latitude, 10)
	if err1 != nil || err2 != nil {
		BuildResponse(context, http.StatusBadRequest, enums.PositionError, "params 'longitude' or 'latitude' error!")
		return
	}
	shopService := service.ShopService{}
	searchShop := shopService.SearchShop(longitudeFloat, latitudeFloat, keyword)
	BuildSuccessResponse(context, searchShop)
}
