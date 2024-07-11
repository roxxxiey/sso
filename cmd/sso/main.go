package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sso/internal/app"
	"sso/internal/config"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TO Do: инициализировать объект конфига

	cfg := config.MustLoad()

	// TO Do: инициализировать логгер

	log := setupLogger(cfg.Env)

	log.Info("starting app", slog.Any("cfg", cfg))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTl)

	go application.GROCSrv.MustRun()

	// TO Do: инициализировать приложение (app)
	// TO Do: запустить grpc сервер приложения

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	application.GROCSrv.Stop()

	log.Info("application Stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
