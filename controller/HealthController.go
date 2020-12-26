package controller

import (
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	BaseController
}

func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/health", hello.Hello)
}

func (hello *HelloController) Hello(context *gin.Context) {
	BuildSuccessResponse(context,"health")
}
