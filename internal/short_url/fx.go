package short_url

import (
	"go.uber.org/fx"
	"tiny-url/internal/short_url/url_generator"
	"tiny-url/internal/storage"
	"tiny-url/internal/url_hash"
)

func Module() fx.Option {
	return fx.Module(
		"short_url",
		fx.Provide(
			NewShortUrl,
		),
		url_generator.Module(),
		storage.Module(),
		url_hash.Module(),
	)
}
