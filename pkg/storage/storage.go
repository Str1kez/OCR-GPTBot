package storage

import (
	"context"
	"strconv"

	"github.com/Str1kez/chatGPT-bot/internal/config"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(config *config.StorageConfig) (*RedisStorage, error) {
	storage := RedisStorage{redis.NewClient(&redis.Options{Addr: config.URL, Password: config.Password, DB: config.DB})}
	if err := storage.client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &storage, nil
}

func (s *RedisStorage) Get(key int64) (string, error) {
	keyStr := strconv.FormatInt(key, 10)
	value, err := s.client.Get(context.Background(), keyStr).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return value, nil
}

func (s *RedisStorage) Set(key int64, value string) error {
	keyStr := strconv.FormatInt(key, 10)
	err := s.client.Set(context.Background(), keyStr, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (s *RedisStorage) Del(key int64) error {
	keyStr := strconv.FormatInt(key, 10)
	err := s.client.Del(context.Background(), keyStr).Err()
	if err != nil {
		return err
	}
	return nil
}
