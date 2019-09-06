package services

import (
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/serializer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteChannelService struct {
}

func (service *DeleteChannelService) Vaild(context *gin.Context) *serializer.Response {
	count := 0

	err := models.DB.Model(&models.Channel{}).Where("id = ?", context.Param("id")).Count(&count).Error

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
			Message: "无法找到对应的记录.",
		}
	}

	return nil
}

func (service *DeleteChannelService) Delete(context *gin.Context) *serializer.Response {
	if err := service.Vaild(context); err != nil {
		return err
	}

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
