package services

import (
	"go-fast-ws-mq/modules"
	"go-fast-ws-mq/serializer"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ListenChannelService struct {
}

func (service *ListenChannelService) Listen(context *gin.Context) *serializer.Response {
	//channelId := context.Param("channelId")
	//connectKey := context.Query("connectKey")

	err := modules.WebSocketModule.HandleRequest(context.Writer, context.Request)

	if err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "升级到Websocket协议时出错.",
			Error:   err.Error(),
		}
	}

	return nil
}
