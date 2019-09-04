package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/serializer"
	"github.com/satori/go.uuid"
)

type CreateChannelService struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

func (service *CreateChannelService) Create() (*models.Channel, *serializer.Response) {
	channel := &models.Channel{
		Name: service.Name,
		Description: service.Description,
		ConnectKey: ,
	}
}
