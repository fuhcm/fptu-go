package redis

import (
	"os"

	"github.com/go-redis/redis"
)

// GetRedisClient ...
func GetRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})
}

// Get ...
func Get(key string) (string, error) {
	client := GetRedisClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

// Set ...
func Set(key string, value string) error {
	client := GetRedisClient()

	err := client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
