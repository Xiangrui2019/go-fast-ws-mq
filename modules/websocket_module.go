package modules

import (
	"log"

	"gopkg.in/olahol/melody.v1"
)

var WebSocketModule *melody.Melody

func InitWebSocketModule() {
	WebSocketModule = melody.New()

	WebSocketModule.HandleConnect(func(session *melody.Session) {
		channelId := session.Request.URL.Query()
		log.Println(channelId)
	})
}
