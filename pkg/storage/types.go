package storage

import (
	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client  *redis.Client
	handler *rejson.Handler
}
