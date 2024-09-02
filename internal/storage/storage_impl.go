package storage

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"log/slog"
)

const (
	putQuery = "INSERT INTO urls (long_url, short_url) VALUES (?, ?)"
	getQuery = "SELECT long_url, short_url FROM urls WHERE short_url = ?"
)

type Params struct {
	fx.In

	DB     *pgxpool.Pool
	Redis  *redis.Client
	Logger slog.Logger
}

type storageImpl struct {
	db     *pgxpool.Pool
	cache  *redis.Client
	logger slog.Logger
}

func NewStorageImpl(params Params) Storage {
	return &storageImpl{
		db:     params.DB,
		cache:  params.Redis,
		logger: params.Logger,
	}
}

func (s storageImpl) Put(ctx context.Context, longURL string, shortURL string) error {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction %w", err)
	}

	_, err = tx.Exec(ctx, putQuery, longURL, shortURL)
	if err != nil {
		rollbackErr := tx.Rollback(ctx)
		if rollbackErr != nil {
			return fmt.Errorf("insert failed: %v, rollback failed: %v", err, rollbackErr)
		}

		s.logger.Debug("rollback", longURL, shortURL)

		return err
	}

	err = s.cache.Set(ctx, longURL, shortURL, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to update cache %w", err)
	}

	return tx.Commit(ctx)
}

func (s storageImpl) Get(ctx context.Context, shortURL string) (string, error) {
	cacheURL, err := s.cache.Get(ctx, shortURL).Result()
	if err != nil {
		return "", err
	}

	if cacheURL != "" {
		return cacheURL, nil
	}

	var url string

	err = s.db.QueryRow(ctx, getQuery, shortURL).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s storageImpl) GetFilter(ctx context.Context) ([]byte, error) {
	var filter []byte

	err := s.db.QueryRow(ctx, "SELECT filter from bloom_filter;").Scan(&filter)
	if err != nil {
		return nil, fmt.Errorf("failed to load actual bloom filter %w", err)
	}

	return filter, nil
}
