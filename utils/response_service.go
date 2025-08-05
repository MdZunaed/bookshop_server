package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseService struct {
	data any
}

func (rs *ResponseService) Success(ctx *gin.Context,statusCode int,message string,  data any) {
	var response = map[string]any{
		"statusCode": statusCode,
		"message": message,
		"data": data,
	}
	if message == ""{
		response["message"] = "success"
	}
	ctx.JSON(statusCode, response)
}

func GetResponseService() *ResponseService{
	return &ResponseService{}
}