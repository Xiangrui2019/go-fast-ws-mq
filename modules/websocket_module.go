package modules

import (
	"gopkg.in/olahol/melody.v1"
)

var WebSocketModule *melody.Melody

func InitWebSocketModule() {
	WebSocketModule = melody.New()

	WebSocketModule.HandleConnect(func(session *melody.Session) {
		channelId := session.Request.URL.Query()["channel_id"][0]
		RedisMQModule.Custome(channelId, func(message string) error {
			session.Write([]byte(message))
			return nil
		})
	})
}
