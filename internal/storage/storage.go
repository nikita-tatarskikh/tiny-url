package storage

import "context"

type Storage interface {
	Put(ctx context.Context, longURL string, shortURL string) error
	Get(ctx context.Context, shortURL string) (string, error)
	GetFilter(ctx context.Context) ([]byte, error)
}
