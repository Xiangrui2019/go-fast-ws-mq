package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/serializer"
	"net/http"
)

type ListChannelService struct {
}

func (service *ListChannelService) List() ([]*models.Channel, *serializer.Response) {
	var channels []*models.Channel

	err := models.DB.Model(&models.Channel{}).Find(&channels).Error

	if err != nil {
		return nil, &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "数据库出错.",
			Error:   err.Error(),
		}
	}

	return channels, nil
}
