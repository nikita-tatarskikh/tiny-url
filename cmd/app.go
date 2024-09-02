package main

import (
	"context"
	"go.uber.org/fx"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"tiny-url/internal/configuration"
	"tiny-url/internal/server"
)

func App() *fx.App {
	app := fx.New(
		configuration.Module(),
		server.Module(),
		fx.Provide(AppCtx),
		fx.Provide(Logger),
		fx.Invoke(func(lifecycle fx.Lifecycle, server server.Server) {
			lifecycle.Append(fx.Hook{
				OnStart: func(context.Context) error {
					return server.Run()
				},
				OnStop: func(ctx context.Context) error {
					return server.Stop()
				},
			})
		},
		),
	)

	return app
}

func AppCtx(lc fx.Lifecycle) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				sigs := make(chan os.Signal, 1)
				signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

				sig := <-sigs
				log.Printf("Received signal: %s. Shutting down...", sig)

				cancel() // Завершаем контекст
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Application is stopping.")
			cancel() // Завершаем контекст, если приложение останавливается
			return nil
		},
	})

	return ctx
}

func Logger() slog.Logger {
	logger := slog.Logger{}

	return logger
}
