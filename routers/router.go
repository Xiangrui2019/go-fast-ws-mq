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
	}

	return router
}
