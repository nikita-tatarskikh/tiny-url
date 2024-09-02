package short_url

import (
	"context"
	"go.uber.org/fx"
	"tiny-url/internal/short_url/url_generator"
	"tiny-url/internal/storage"
	"tiny-url/internal/url_hash"
)

type ShortenerParams struct {
	fx.In

	Storage   storage.Storage
	UrlHash   url_hash.URLHash
	Generator url_generator.Generator
}

type shortUrlImpl struct {
	storage   storage.Storage
	urlHash   url_hash.URLHash
	generator url_generator.Generator
}

func NewShortUrl(params ShortenerParams) ShortUrl {
	return &shortUrlImpl{
		storage:   params.Storage,
		urlHash:   params.UrlHash,
		generator: params.Generator,
	}
}

func (s shortUrlImpl) MakeShortURL(ctx context.Context, url string) (string, error) {
	shortURL := s.generator.GenerateShortURL(url)

	err := s.storage.Put(ctx, url, shortURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}
