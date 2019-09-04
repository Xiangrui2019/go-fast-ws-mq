package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/serializer"
)

type CreateChannelService struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

func (service *CreateChannelService) Create() (*models.Channel, *serializer.Response) {

}
