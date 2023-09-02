package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/nitishm/go-rejson/v4"

	"github.com/Str1kez/OCR-GPTBot/internal/config"
	"github.com/Str1kez/OCR-GPTBot/pkg/telegram"
	"github.com/nitishm/go-rejson/v4/rjs"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func NewRedisStorage(config *config.StorageConfig) (*RedisStorage, error) {
	redisClient := redis.NewClient(&redis.Options{Addr: config.URL, Password: config.Password, DB: config.DB})
	rejsonHandler := rejson.NewReJSONHandler()
	rejsonHandler.SetGoRedisClient(redisClient)
	storage := RedisStorage{client: redisClient, handler: rejsonHandler}
	if err := storage.client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &storage, nil
}

func (s *RedisStorage) Get(userId int64, key string) ([]byte, error) {
	documentName := fmt.Sprintf("user_settings:%d", userId)
	result, err := s.handler.JSONGet(documentName, "."+key)
	if err != nil {
		if errors.Is(err, redis.Nil) || strings.Contains(err.Error(), "does not exist") {
			return []byte{}, nil
		}
		log.Errorf("Couldn't get user settings with key. Document: %s | Key = %s\n%v\n", documentName, key, err)
		return []byte{}, err
	}
	resultBytes, ok := result.([]byte)
	if !ok {
		log.Errorf("Couldn't convert response to slice of bytes")
		return []byte{}, &ParseError{}
	}
	return resultBytes, nil
}

func (s *RedisStorage) GetAll(userId int64) (telegram.Settings, error) {
	documentName := fmt.Sprintf("user_settings:%d", userId)
	data, err := s.handler.JSONGet(documentName, ".")
	if err != nil {
		if errors.Is(err, redis.Nil) {
			settings := telegram.GetDefaultSettings()
			err = s.SetAll(userId, settings)
			if err != nil {
				return telegram.Settings{}, err
			}
			return settings, nil
		}
		log.Errorf("Couldn't get user settings. Document: %s\n%v\n", documentName, err)
		return telegram.Settings{}, err
	}
	bytesData, ok := data.([]byte)
	if !ok {
		log.Errorln("Couldn't represent data to slice of bytes")
		return telegram.Settings{}, &ParseError{}
	}
	var result telegram.Settings
	err = json.Unmarshal(bytesData, &result)
	if err != nil {
		log.Errorln("Couldn't parse data to settings struct")
		return telegram.Settings{}, &ParseError{}
	}
	return result, nil
}

func (s *RedisStorage) Set(userId int64, key string, value interface{}) error {
	documentName := fmt.Sprintf("user_settings:%d", userId)
	result, err := s.handler.JSONSet(documentName, "."+key, value)
	if err != nil && strings.Contains(err.Error(), "new objects must be created at the root") {
		settings := telegram.GetDefaultSettings()
		err = s.SetAll(userId, settings)
		if err != nil {
			return err
		}
		result, err = s.handler.JSONSet(documentName, "."+key, value)
	}
	if err != nil {
		log.Errorf("Couldn't set data. Document: %s | Key: %s\n%v\n", documentName, key, err)
		return err
	}
	if status, ok := result.(string); ok && status != "OK" {
		log.Errorf("Returned status after save: %s | Document: %s | Key : %s\n", status, documentName, key)
		return &SetError{}
	}
	log.Debugln("Data saved")
	return nil
}

func (s *RedisStorage) SetAll(userId int64, settings telegram.Settings) error {
	documentName := fmt.Sprintf("user_settings:%d", userId)
	result, err := s.handler.JSONSet(documentName, ".", settings, rjs.SetOptionNX)
	if err != nil {
		log.Errorf("Couldn't set data. Document: %s\n%v\n", documentName, err)
		return err
	}
	if status, ok := result.(string); ok && status != "OK" {
		log.Errorf("Returned status after save: %s | Document: %s\n", status, documentName)
		return &SetError{}
	}
	log.Debugln("Data saved")
	return nil
}

func (s *RedisStorage) Del(userId int64, key string) error {
	documentName := fmt.Sprintf("user_settings:%d", userId)
	result, err := s.handler.JSONDel(documentName, "."+key)
	if err != nil {
		log.Errorf("Couldn't delete data. Document: %s | Key: %s\n%v\n", documentName, key, err)
		return err
	}
	if status, ok := result.(string); ok && status != "OK" {
		log.Errorf("Returned status after delete: %s | Document: %s | Key: %s\n", status, documentName, key)
		return &SetError{}
	}
	log.Debugln("Data removed")
	return nil
}

func (s *RedisStorage) DelAll(userId int64) error {
	documentName := fmt.Sprintf("user_settings:%d", userId)
	err := s.client.Del(context.Background(), documentName).Err()
	if err != nil {
		log.Errorf("Couldn't delete data. Document: %s\n%v\n", documentName, err)
	}
	log.Debugln("Data removed")
	return err
}
