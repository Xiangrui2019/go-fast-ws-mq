package serializer

import "go-fast-ws-mq/models"

type PrivateChannel struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ConnectKey  string `json:"connect_key"`
	CreatedAt   int64  `json:"created_at"`
}

type PublicChannel struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
}

// BuildPrivateChannel 序列化私有Channel
func BuildPrivateChannel(channel *models.Channel) PrivateChannel {
	return PrivateChannel{
		ID:          channel.ID,
		Name:        channel.Name,
		Description: channel.Description,
		ConnectKey:  channel.ConnectKey,
		CreatedAt:   channel.CreatedAt.Unix(),
	}
}

// BuildPublicChannel 序列化公开Channel
func BuildPublicChannel(channel *models.Channel) PublicChannel {
	return PublicChannel{
		ID:          channel.ID,
		Name:        channel.Name,
		Description: channel.Description,
		CreatedAt:   channel.CreatedAt.Unix(),
	}
}
