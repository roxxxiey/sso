package app

import (
	"log/slog"
	grpcapp "sso/internal/app/grpc"
	"time"
)

type App struct {
	GROCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TO DO: инициализировать хранилище (storage)

	// init auth service (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GROCSrv: grpcApp,
	}
}
