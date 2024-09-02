package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"tiny-url/internal/configuration"
	"tiny-url/pkg"
)

type Params struct {
	fx.In

	AppCtx context.Context
	Config configuration.Config
}

func NewConnPool(lc fx.Lifecycle, params Params) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf(
		pkg.PostgresqlConnectionString,
		params.Config.PostgresConfig.PostgresUser,
		params.Config.PostgresConfig.PostgresPassword,
		params.Config.PostgresConfig.PostgresAddress,
		params.Config.PostgresConfig.PostgresPort,
		params.Config.PostgresConfig.PostgresDatabase,
	)

	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse parse database config: %w", err)
	}

	db, err := pgxpool.NewWithConfig(params.AppCtx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return fmt.Errorf("unable to verify postgres connection %w", db.Ping(ctx))
		},
		OnStop: func(ctx context.Context) error {
			db.Close()
			return nil
		},
	})

	return db, nil
}
