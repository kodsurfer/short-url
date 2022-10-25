package redis

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewClient(n int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASS"),
		DB:       n,
	})
}
