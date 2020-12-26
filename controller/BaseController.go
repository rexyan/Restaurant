package controller

import (
	"Restaurant/enums"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

//BuildResponse
func BuildResponse(context *gin.Context, HttpStatus int, MessageCode string, Data interface{}) {
	requestId, _ := context.Get("requestId")
	context.JSON(HttpStatus, gin.H{
		"data":        Data,
		"message":     enums.ErrorMessage[MessageCode],
		"return_code": MessageCode,
		"response_id": requestId,
	})
}

// BuildSuccessResponse
func BuildSuccessResponse(context *gin.Context, Data interface{}) {
	BuildResponse(context, http.StatusOK, enums.SUCCESS, Data)
}
