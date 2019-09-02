package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/modules"
	"go-fast-ws-mq/serializer"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ListenChannelService struct {
	ChannelId  int `form:"channel_id" json:"channel_id" binding:"required"`
	ConnectKey int `form:"connect_key" json:"connect_key" binding:"required"`
}

func (service *ListenChannelService) Vaild() *serializer.Response {
	count := 0
	err := models.DB.Model(&models.Channel{}).Where("id = ?", service.ChannelId).Count(&count).Error

	if err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "数据库出错.",
			Error:   err.Error(),
		}
	}

	if count == 0 {
		return &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "无法找到对应的频道.",
		}
	}

	err = models.DB.Model(&models.Channel{}).Where("x = ?", service.ConnectKey).Count(&count).Error

	if err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "数据库出错.",
			Error:   err.Error(),
		}
	}

	if count == 0 {
		return &serializer.Response{
			Code:    http.StatusBadRequest,
			Message: "ConnectKey 有误.",
		}
	}

	return nil
}

func (service *ListenChannelService) Listen(context *gin.Context) *serializer.Response {
	if err := service.Vaild(); err != nil {
		return err
	}

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
