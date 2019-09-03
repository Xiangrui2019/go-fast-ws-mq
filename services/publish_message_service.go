package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/modules"
	"go-fast-ws-mq/serializer"
	"net/http"
	"strconv"
)

type PublishMessageService struct {
	ChannelId int    `form:"channel_id" json:"channel_id" binding:"required"`
	Message   string `form:"message" json:"message" binding:"required"`
}

func (service *PublishMessageService) Vaild() *serializer.Response {
	count := 0
	err := models.DB.Model(&models.Channel{}).Where("id", service.ChannelId).Count(&count).Error

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

	return nil
}

func (service *PublishMessageService) Publish() *serializer.Response {
	if err := service.Vaild(); err != nil {
		return err
	}

	err := modules.RedisMQModule.Publish(strconv.Itoa(service.ChannelId), service.Message)

	if err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "发布消息失败.",
			Error:   err.Error(),
		}
	}

	return &serializer.Response{
		Code:    http.StatusOK,
		Message: "发布消息成功.",
	}
}
