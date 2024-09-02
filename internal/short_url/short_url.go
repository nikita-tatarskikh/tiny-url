package short_url

import (
	"context"
)

type ShortUrl interface {
	MakeShortURL(ctx context.Context, url string) (string, error)
}
