package controller

import (
	"Restaurant/enums"
	"Restaurant/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GoodsController struct {
	BaseController
}

func (gc *GoodsController) Router(engine *gin.Engine) {
	engine.GET("/api/foods", gc.getGoodsByShopId)
}

// getGoods 查询某个商户的商品
func (gc *GoodsController) getGoodsByShopId(context *gin.Context) {
	shopId, exist := context.GetQuery("shopId")
	id, err := strconv.Atoi(shopId)
	if err != nil || !exist {
		BuildResponse(context, http.StatusBadRequest, enums.ParamError, "params 'shopId' error!")
		return
	}
	goodsService := service.GoodsService{}
	goods := goodsService.GetGoodsByShopId(int64(id))
	BuildSuccessResponse(context, goods)
}
