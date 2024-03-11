package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	authorization "url-shortener"
	"url-shortener/internal/config"
	"url-shortener/pkg/handler"
	"url-shortener/pkg/repository"
	"url-shortener/pkg/service"

	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	cfg, err := config.MustLoad(os.Getenv("CONFIG_PATH"), os.Getenv("CONFIG_NAME"), os.Getenv("CONFIG_TYPE"))
	if err != nil {
		fmt.Printf("error %s", err)
	}

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages enabled")

	storage, err := repository.NewSQLStorage(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage ", err)
		os.Exit(1)
	}

	repository := repository.NewRepository(storage)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(authorization.Server)

	if err := server.Run(cfg.Address, cfg.Timeout, cfg.IdleTimeout, handlers.InitRoutes()); err != nil {
		log.Error("Failed to run http server: ", err)
		os.Exit(1)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
