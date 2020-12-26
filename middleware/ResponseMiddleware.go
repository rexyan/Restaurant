package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func ResponseMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := uuid.NewV4()
		context.Set("requestId", requestId)
		context.Next()
	}
}