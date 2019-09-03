package services

import "go-fast-ws-mq/serializer"

type PublishMessageService struct {
	ChannelId  int
	ConnectKey int
	Message    string
}

func (service *PublishMessageService) Publish() *serializer.Response {
	return nil
}
