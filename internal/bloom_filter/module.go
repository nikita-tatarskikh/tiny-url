package bloom_filter

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("bloom_filter", fx.Provide(New))
}
