package storage

import (
	"go.uber.org/fx"
	"tiny-url/pkg/postgres"
	"tiny-url/pkg/redis"
)

func Module() fx.Option {
	return fx.Module(
		"storage",
		fx.Provide(NewStorageImpl),
		postgres.Module(),
		redis.Module(),
	)
}
