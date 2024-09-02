package server

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
	"go.uber.org/fx"
	"log/slog"
	"tiny-url/internal/configuration"
)

type Params struct {
	fx.In

	FastHttpServer *fasthttp.Server
	Config         configuration.Config
	AppCtx         context.Context
	Logger         slog.Logger
}

type Server struct {
	FastHttpServer *fasthttp.Server
	Config         configuration.Config
	Ctx            context.Context
	Logger         slog.Logger
}

func NewServer(params Params) Server {
	return Server{
		FastHttpServer: params.FastHttpServer,
		Config:         params.Config,
		Ctx:            params.AppCtx,
	}
}

func (s Server) Run() error {
	addr := fmt.Sprintf("%s:%s", s.Config.ServerConfig.ServerAddress, s.Config.ServerConfig.ServerPort)

	s.Logger.Info("server started")

	return s.FastHttpServer.ListenAndServe(addr)
}

func (s Server) Stop() error {
	return s.FastHttpServer.Shutdown()
}
