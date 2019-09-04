package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/serializer"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type CreateChannelService struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}

func (service *CreateChannelService) Vaild() *serializer.Response {
	count := 0
	err := models.DB.Model(&models.Channel{}).Where("name = ?", service.Name).Count(&count).Error

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

func (service *CreateChannelService) Create() (*models.Channel, *serializer.Response) {
	if err := service.Vaild(); err != nil {
		return nil, err
	}

	channel := &models.Channel{
		Name:        service.Name,
		Description: service.Description,
		ConnectKey:  uuid.NewV4().String(),
	}

	result := models.DB.Create(&channel)

	if result.Error != nil {
		return nil, &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "数据库创建错误.",
			Error:   result.Error.Error(),
		}
	}

	return result.Value.(*models.Channel), nil
}
