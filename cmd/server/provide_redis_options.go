package main

import (
	"github.com/go-redis/redis/v8"
)

func ProvideRedisOptions() *redis.Options {
	return &redis.Options{
		Network: "tcp",
		Addr:    redisUrl,
		DB:      redisDb,
	}
}
