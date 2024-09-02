package fast_http_router

import (
	"go.uber.org/fx"
	"tiny-url/internal/usecase/make_short_url"
)

func Module() fx.Option {
	return fx.Module(
		"router", fx.Provide(NewRouter),
		make_short_url.Module(),
	)
}
