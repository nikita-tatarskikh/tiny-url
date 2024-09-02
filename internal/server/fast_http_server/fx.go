package fast_http_server

import (
	"go.uber.org/fx"
	"tiny-url/internal/fast_http_router"
)

func FastHttpServerModule() fx.Option {
	return fx.Module(
		"fast_http_server",
		fx.Provide(NewFastHttpServer),
		fast_http_router.Module(),
	)
}
