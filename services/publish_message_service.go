package services

import (
	"go-fast-ws-mq/modules"
	"go-fast-ws-mq/serializer"
	"strconv"
)

type PublishMessageService struct {
	ChannelId  int
	ConnectKey int
	Message    string
}

func (service *PublishMessageService) Publish() *serializer.Response {
	err := modules.RedisMQModule.Publish(strconv.Itoa(service.ChannelId), service.Message)

	if err != nil {
		return &serializer.Response{}
	}

	return &serializer.Response{}
}
