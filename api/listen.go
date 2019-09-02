package api

import (
	"go-fast-ws-mq/serializer"
	"go-fast-ws-mq/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListenChannel(context *gin.Context) {
	service := services.ListenChannelService{}

	if err := context.ShouldBindQuery(&service); err != nil {
		context.JSON(http.StatusBadRequest, &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "您的输入有误.",
			Error:   err.Error(),
		})
	} else {
		if err := service.Listen(context); err != nil {
			context.JSON(err.Code, err)
		}
	}
}
