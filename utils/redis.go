package utils

import (
	"app/config"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type redisUtils struct {
	redisClient *redis.Client
}

type RedisUtils interface {
	Cache(key string, value interface{}) error
	Delete(key string) error
	GetData(key string) (interface{}, error)
}

func (u *redisUtils) Cache(key string, value interface{}) error {
	dataByte, _ := json.Marshal(value)
	err := u.redisClient.Set(context.Background(), key, dataByte, 0).Err()
	return err
}

func (u *redisUtils) Delete(key string) error {
	err := u.redisClient.Del(context.Background(), key).Err()
	return err
}

func (u *redisUtils) GetData(key string) (interface{}, error) {
	var value interface{}
	result, err := u.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(result), &value); err != nil {
		return nil, err
	}

	return value, nil
}

func NewUtilsRedis() RedisUtils {
	return &redisUtils{
		redisClient: config.GetRDB(),
	}
}
