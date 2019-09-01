package conf

import (
	"go-fast-ws-mq/cache"
	"go-fast-ws-mq/models"
	"go-fast-ws-mq/modules"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	models.ConnectDatabase(os.Getenv("DATABASE_DSN"))
	cache.ConnectRedisCache()
	modules.InitAllModules()
}
