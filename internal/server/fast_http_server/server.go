package fast_http_server

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"go.uber.org/fx"
)

type FastHTTPServerParams struct {
	fx.In

	Router *router.Router
}

func NewFastHttpServer(params FastHTTPServerParams) *fasthttp.Server {
	return &fasthttp.Server{
		Handler: params.Router.Handler,
	}
}
