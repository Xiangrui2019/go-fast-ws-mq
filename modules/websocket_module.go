package modules

import (
	"gopkg.in/olahol/melody.v1"
)

var WebSocketModule *melody.Melody

func InitWebSocketModule() {
	WebSocketModule = melody.New()

	WebSocketModule.HandleConnect(func(session *melody.Session) {
		channelId := session.Request.URL.Query()["channelId"][0]
		RedisMQModule.Custome(channelId, func(message string) error {
			return nil
		})
	})
}