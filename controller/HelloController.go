package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloController struct {
}

func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/health", hello.Hello)
}

func (hello *HelloController) Hello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hi Restaurant!",
	})
}
