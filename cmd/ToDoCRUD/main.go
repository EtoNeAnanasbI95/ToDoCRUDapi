// @title ToDo CRUD API
// @version 1.0
// @description API для управления задачами и пользователями
// @host localhost:8080
// @BasePath /api
package main

import (
	"fmt"
	"github.com/EtoNeAnanasbI95/ToDoCRUD"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/config"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/handler"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/repository"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/service"
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/storage"
	"log/slog"
	"os"
	"sync"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoadConfig()

	log := setupLogger(cfg.Env)

	log.Info("Starting CRUD api",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.Api.Port))

	db := storage.MustInitDB(cfg.ConnectionString)
	defer db.Close()
	r := repository.NewRepository(db)
	//TODO: инициализировать сервисы
	s := service.NewService(r)
	// TODO: инициализировать хендлеры
	handler := handler.NewHandler(log, s)
	api := handler.InitRouts()
	// TODO: запустить апи
	srv := new(ToDoCRUD.Server)
	wg := sync.WaitGroup{}
	wg.Add(1)
	//TODO: поправить выдачу статус кодов
	go func() {
		if err := srv.Run(api, cfg); err != nil {
			log.Error(err.Error())
		}
		wg.Done()
	}()
	log.Info("Api is running")
	if cfg.Env == envLocal {
		log.Info("Running in local mode",
			slog.String("URL", fmt.Sprintf("http://localhost:%d/swagger/index.html", cfg.Api.Port)))
	}
	// TODO: сделать gracefull shutdown
	wg.Wait()
	// TODO: написать докер
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger
}
