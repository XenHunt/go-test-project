package main

import (
	"log/slog"
	"os"

	"github.com/XenHunt/go-test-project/internal/config"
	db_module "github.com/XenHunt/go-test-project/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger()
	log = log.With("env", cfg.Env)

	// log.Info("", args ...any)
	log.Debug("Debug started")
	log.Debug("%s", cfg)
	// fmt.Println(cfg.DataBaseConfig)
	database := db_module.MakeConection(cfg.DataBaseConfig)
	router := chi.NewRouter()

	router.Use(middleware.Logger)
}

func setupLogger() *slog.Logger {
	//TODO надо добавить различные уровни логирования
	var log *slog.Logger
	log = slog.New(slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: slog.LevelDebug},
	))
	return log
}
