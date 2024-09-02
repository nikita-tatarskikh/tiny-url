package make_short_url

import (
	"go.uber.org/fx"
	"tiny-url/internal/short_url"
)

func Module() fx.Option {
	return fx.Module(
		"make_short_url",
		fx.Provide(NewShortURLHandler),
		short_url.Module(),
	)
}
