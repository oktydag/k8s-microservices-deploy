package repository

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func InitializeRedis() (string) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis", // set docker-compose service name
	})

	ping, err := redisClient.Ping().Result()
	var message = ""
	if err == nil && len(ping) > 0 {
		message= "Connected to Redis"
	} else {
		message= "Redis Connection Failed"
	}

	return message
}

func GetValue(key string) (string, error) {
	var deserializedValue interface{}
	serializedValue, err := redisClient.Get(key).Result()
	json.Unmarshal([]byte(serializedValue), &deserializedValue)
	return fmt.Sprintf("%v",deserializedValue), err
}

func SetValue(key string, value interface{}) (bool, error) {
	serializedValue, _ := json.Marshal(value)
	err := redisClient.Set(key, string(serializedValue), 0).Err()
	return true, err
}

func IncrementValue(key string) (int64)  {
	return redisClient.Incr(key).Val()
}
