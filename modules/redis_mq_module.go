package modules

import (
	"go-fast-ws-mq/cache"
	"go-fast-ws-mq/global"
	"log"
)

type RedisMQ struct {
}

type MessageCallback func(message string) error

var RedisMQModule *RedisMQ

func (mq *RedisMQ) Publish(queuename string, message string) error {
	err := cache.CacheClient.Publish(global.QueueNameKey(queuename), message).Err()

	if err != nil {
		return err
	}

	return nil
}

func (mq *RedisMQ) Custome(queuename string, cb MessageCallback) {
	sub := cache.CacheClient.Subscribe(global.QueueNameKey(queuename))

	for message := range sub.Channel() {
		err := cb(message.Payload)

		if err != nil {
			log.Printf("Execute Callback func Error: %s", err)
		}
	}
}

func InitRedisMQModule() {
	RedisMQModule = new(RedisMQ)
}
