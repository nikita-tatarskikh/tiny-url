package server

import (
	"go.uber.org/fx"
	"tiny-url/internal/server/fast_http_server"
)

func Module() fx.Option {
	return fx.Module(
		"server",
		fx.Provide(NewServer),
		fast_http_server.FastHttpServerModule(),
	)
}
