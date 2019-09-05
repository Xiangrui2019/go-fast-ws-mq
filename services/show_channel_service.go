package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowChannelService struct {
}

func (service *ShowChannelService) Show(context *gin.Context) (*models.Channel, *serializer.Response) {
	var channel models.Channel

	err := models.DB.Model(&models.Channel{}).Where("id = ?", context.Param("id")).First(&channel).Error

	if err != nil {
		return nil, &serializer.Response{
			Code:    http.StatusNotFound,
			Message: "找不到对应的ID.",
		}
	}

	return &channel, nil
}
