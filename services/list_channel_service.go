package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/serializer"
)

type ListChannelService struct {
}

func (service *ListChannelService) List() ([]*models.Channel, *serializer.Response) {
	return make([]*models.Channel, 0), nil
}
