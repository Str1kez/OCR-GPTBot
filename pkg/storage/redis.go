package storage

import (
	ctx "context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nitishm/go-rejson/v4"

	"github.com/Str1kez/OCR-GPTBot/internal/config"
	"github.com/Str1kez/OCR-GPTBot/pkg/telegram"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func getDocumentName(userId int64) string {
	return fmt.Sprintf("user_settings:%d", userId)
}

func NewRedisStorage(config *config.StorageConfig) (*RedisStorage, error) {
	redisClient := redis.NewClient(&redis.Options{Addr: config.URL, Password: config.Password, DB: config.DB})
	rejsonHandler := rejson.NewReJSONHandler()
	rejsonHandler.SetGoRedisClientWithContext(ctx.Background(), redisClient)
	storage := RedisStorage{client: redisClient, handler: rejsonHandler}
	if err := storage.client.Ping(ctx.Background()).Err(); err != nil {
		return nil, err
	}
	return &storage, nil
}

func (s *RedisStorage) Get(userId int64) (telegram.Settings, error) {
	documentName := getDocumentName(userId)
	settings := telegram.Settings{}
	data, err := s.handler.JSONGet(documentName, ".")
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return settings, telegram.ErrNotFound
		}
		log.Errorf("Couldn't get user settings. Document: %s\n%v\n", documentName, err)
		return settings, err
	}
	bytesData, ok := data.([]byte)
	if !ok {
		log.Errorln("Couldn't represent data to slice of bytes")
		return settings, ParseError
	}
	err = json.Unmarshal(bytesData, &settings)
	if err != nil {
		log.Errorln("Couldn't parse data to settings struct")
		return settings, ParseError
	}
	return settings, nil
}

func (s *RedisStorage) Set(userId int64, settings telegram.Settings) error {
	documentName := getDocumentName(userId)
	result, err := s.handler.JSONSet(documentName, ".", settings)
	if err != nil {
		log.Errorf("Couldn't set data. Document: %s\n%v\n", documentName, err)
		return err
	}
	if status, ok := result.(string); ok && status != "OK" {
		log.Errorf("Returned status after save: %s | Document: %s\n", status, documentName)
		return SetError
	} else if !ok {
		log.Errorf("Error in converting status after set | Document: %s\n", documentName)
		return SetError
	}
	log.Debugln("Data saved")
	return nil
}

func (s *RedisStorage) Close() error {
	return s.client.Close()
}
