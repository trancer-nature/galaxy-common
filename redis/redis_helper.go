package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(host, password string, selected int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       selected,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
		return nil, err
	}
	return client, nil
}
