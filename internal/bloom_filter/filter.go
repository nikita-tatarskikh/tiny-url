package bloom_filter

import (
	"context"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"go.uber.org/fx"
	"tiny-url/internal/storage"
)

type FilterParams struct {
	fx.In

	Ctx     context.Context
	Storage storage.Storage
}

func New(ctx context.Context, storage storage.Storage) (*bloom.BloomFilter, error) {
	filterBytes, err := storage.GetFilter(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to setup bloom filter: %w", err)
	}

	if filterBytes == nil {
		return bloom.NewWithEstimates(365_000_000_000, 0.01), nil
	}

	filter := bloom.New(0, 0)
	err = filter.UnmarshalJSON(filterBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to setup bloom filter %w", err)
	}

	return filter, nil
}
