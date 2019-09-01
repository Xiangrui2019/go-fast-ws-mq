package conf

import (
	"go-fast-ws-mq/cache"
	"go-fast-ws-mq/models"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	models.ConnectDatabase(os.Getenv("DATABASE_DSN"))
	cache.ConnectRedisCache()
}
