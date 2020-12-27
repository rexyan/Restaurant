package controller

import (
	"Restaurant/enums"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
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

// RequestData, Unmarshal request data
func RequestData(context *gin.Context) map[string]string {
	requestData := map[string]string{}
	data, _ := ioutil.ReadAll(context.Request.Body)
	if err := json.Unmarshal(data, &requestData); err != nil {
		log.Println("params unmarshal error!")
		return nil
	}
	return requestData
}
