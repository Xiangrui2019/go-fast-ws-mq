package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteChannelService struct {
}

func (service *DeleteChannelService) Delete(context *gin.Context) *serializer.Response {
	err := models.DB.Model(&models.Channel{}).Where("id = ?", context.Param("id")).Delete(&models.Channel{}).Error

	if err != nil {
		return &serializer.Response{
			Code:    http.StatusInternalServerError,
			Message: "删除失败",
			Error:   err.Error(),
		}
	}

	return &serializer.Response{
		Code:    http.StatusOK,
		Message: "删除Channel成功.",
	}
}
