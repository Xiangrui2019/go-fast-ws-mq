package routers

import (
	"go-fast-ws-mq/api"
	"go-fast-ws-mq/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// 中间件, 顺序必须是这样
	router.Use(middlewares.Cors(os.Getenv("CORS_DOMAIN")))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/ping", api.Ping)

		// 消息接口
		v1.POST("/message/publish", api.PublishMessage)
		v1.GET("/channel/listen", api.ListenChannel)

		// 频道CRUD接口
		v1.POST("/channel", api.CreateChannel)
	}

	return router
}
