package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
	"tiny-url/internal/configuration"
)

type Params struct {
	fx.In

	AppCtx context.Context
	Config configuration.Config
}

func NewRedisConnection(lc fx.Lifecycle, params Params) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			params.Config.RedisConfig.RedisAddress,
			params.Config.RedisConfig.RedisPort,
		),
		Password: params.Config.RedisConfig.RedisPassword,
		DB:       params.Config.RedisConfig.RedisDB,
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return fmt.Errorf("fail to verify redis connection %w", rdb.Ping(params.AppCtx).Err())
		},
		OnStop: func(ctx context.Context) error {
			return fmt.Errorf("failed to close connection %w", rdb.Close())
		},
	})

	return rdb, nil
}
