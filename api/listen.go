package api

import (
	"go-fast-ws-mq/services"

	"github.com/gin-gonic/gin"
)

func ListenChannel(context *gin.Context) {
	service := services.ListenChannelService{}

	if err := service.Listen(context); err != nil {
		context.JSON(err.Code, err)
	}
}
