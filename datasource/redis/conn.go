package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Connect(uri string) *redis.Client {
	opt, err := redis.ParseURL(uri)
	if err != nil {
		zap.L().Panic("Failed to parse URL", zap.String("URL", uri))
	}

	rdb := redis.NewClient(opt)
	if rdb.Ping(context.TODO()).Err() != nil {
		zap.L().Panic("Failed to Ping", zap.Error(err))
	}
	return rdb
}
