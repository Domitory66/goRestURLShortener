package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/pkg/repository"
	"url-shortener/pkg/storage"

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

	storage, err := storage.NewSQLStorage(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to init storage ", err)
		os.Exit(1)
	}

	repository := repository.NewRepository(storage)

	err = repository.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to save", err)
	}

	url, err := repository.GetURL("google")
	if err != nil {
		log.Error("failed to get", err)
	}
	fmt.Println(url)

	err = repository.DeleteURL("google")
	if err != nil {
		log.Error("failed to delete", err)
	}
	//TODO init router: gin

	//TODO init server
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
