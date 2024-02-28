package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/davewhit3/std-server/ctrl"
	"github.com/davewhit3/std-server/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	slogformatter "github.com/samber/slog-formatter"
)

var (
	AppEnv  = "production"
	AppPort = 8080

	cfg = server.Config{
		Env:  AppEnv,
		Port: AppPort,
	}
)

func main() {
	env.Parse(&cfg)

	ctx := context.Background()

	logger := slog.New(
		slogformatter.NewFormatterHandler(
			slogformatter.ErrorFormatter("error"),
		)(
			slog.NewTextHandler(os.Stdout, nil).WithAttrs([]slog.Attr{slog.String("env", cfg.Env)}),
		),
	)

	cfg.Logger = logger

	srv := server.New(cfg)

	srv.RegisterController(ctrl.SetupUserHandlers)
	srv.RegisterController(ctrl.SetupHealthHandlers)
	srv.RegisterHandler(promhttp.Handler())

	if err := srv.Run(ctx); err != nil {
		logger.Error("cannot run server", slog.Any("error", err))
	}
}
