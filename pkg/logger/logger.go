// Package logger -
package logger

import (
	"fmt"
	"log/slog"
	"os"
)

const (
	envLoc  = "local"
	envDev  = "dev"
	envProd = "prod"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	file, err := os.OpenFile("log/logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		panic(fmt.Errorf("не удача открытия файла: %w", err))
	}

	switch env {
	case envLoc:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(file, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
