package api

import (
	"go-fast-ws-mq/serializer"
	"go-fast-ws-mq/services"
	"go-fast-ws-mq/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateChannel(context *gin.Context) {
	service := services.CreateChannelService{}

	if err := context.ShouldBind(&service); err == nil {
		if channel, err := service.Create(); err != nil {
			context.JSON(err.Code, err)
		} else {
			context.JSON(http.StatusOK,
				serializer.BuildPrivateChannelResponse(channel))
		}
	} else {
		context.JSON(http.StatusBadRequest, utils.BuildErrorResponse(err))
	}
}
