package fast_http_router

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
	"log/slog"
)

type Handlers struct {
	fx.In

	ShortURLHandler fasthttp.RequestHandler
	LongURLHandler  fasthttp.RequestHandler
	Logger          slog.Logger
}

func NewRouter(handlers Handlers) *router.Router {
	rtr := router.New()

	rtr.POST("/", handlers.ShortURLHandler)
	rtr.GET("/{short-url}", handlers.LongURLHandler)

	handlers.Logger.Info("registered", rtr.List())

	return rtr
}
