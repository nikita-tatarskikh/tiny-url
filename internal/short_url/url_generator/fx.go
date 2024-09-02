package url_generator

import (
	"go.uber.org/fx"
	"tiny-url/internal/bloom_filter"
)

func Module() fx.Option {
	return fx.Module(
		"url_generator",
		fx.Provide(NewGenerator),
		bloom_filter.Module(),
	)
}
