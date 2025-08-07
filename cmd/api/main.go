package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/charmingruby/pack/config"
	"github.com/charmingruby/pack/internal/platform"
	"github.com/charmingruby/pack/pkg/database/postgres"
	"github.com/charmingruby/pack/pkg/delivery/http/rest"
	"github.com/charmingruby/pack/pkg/telemetry/logger"

	"github.com/joho/godotenv"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Warn("failed to find .env file", "error", err)
	}

	log.Info("loading environment variables...")

	cfg, err := config.New()
	if err != nil {
		log.Error("failed to loading environment variables", "error", err)
		failAndExit(log, nil, nil)
	}

	log.Info("environment variables loaded")

	logLevel := logger.ChangeLevel(cfg.LogLevel)

	log.Info("log level configured", "level", logLevel)

	db, err := postgres.New(log, cfg.PostgresURL)
	if err != nil {
		log.Error("failed to loading environment variables", "error", err)
		failAndExit(log, nil, nil)
	}

	srv, r := rest.New(cfg.RestServerPort)

	platform.New(r, db)

	go func() {
		log.Info("REST server is running...", "port", cfg.RestServerPort)

		if err := srv.Start(); err != nil {
			log.Error("failed starting REST server", "error", err)
			failAndExit(log, srv, db)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info("received an interrupt signal")

	log.Info("starting graceful shutdown...")

	signal := gracefulShutdown(log, srv, db)

	log.Info(fmt.Sprintf("gracefully shutdown, with code %d", signal))

	os.Exit(signal)
}

func failAndExit(log *logger.Logger, srv *rest.Server, db *postgres.Client) {
	gracefulShutdown(log, srv, db)
	os.Exit(1)
}

func gracefulShutdown(log *logger.Logger, srv *rest.Server, db *postgres.Client) int {
	parentCtx := context.Background()

	var hasError bool

	if srv != nil {
		ctx, cancel := context.WithTimeout(parentCtx, 15*time.Second)
		defer cancel()

		if err := srv.Stop(ctx); err != nil {
			log.Error("error closing REST server", "error", err)
			hasError = true
		}
	}

	if db != nil {
		if err := db.Close(); err != nil {
			log.Error("error closing Postgres connection", "error", err)
			hasError = true
		}
	}

	if hasError {
		return 1
	}

	return 0
}
