package api

import (
	"go-fast-ws-mq/services"
	"go-fast-ws-mq/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublishMessage(context *gin.Context) {
	service := services.PublishMessageService{}

	if err := context.ShouldBind(&service); err == nil {
		result := service.Publish()

		context.JSON(result.Code, result)
	} else {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	}
}
