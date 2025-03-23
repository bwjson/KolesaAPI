package main

import (
	_ "github.com/bwjson/api/docs"
	"github.com/bwjson/api/internal"
	"github.com/bwjson/api/internal/config"
	"github.com/bwjson/api/internal/postgres"
	"github.com/bwjson/api/internal/repository"
	"github.com/bwjson/api/internal/service"
	"github.com/bwjson/api/internal/transport"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

// @title           Kolesa API
// @version         1.0
// @description     This is a sample server celler server.
// @host      localhost:8000
// @BasePath  /api
func main() {
	cfg := config.LoadConfig()

	log := setupLogger(cfg.Env)
	log.Info("Start the application", slog.String("address", cfg.Address))

	db, err := postgres.NewPostgresDB(cfg.Db.User, cfg.Db.Name, cfg.Db.Port, cfg.Db.Password, cfg.Db.Host, cfg.Db.Sslmode)
	if err != nil {
		log.Error("Cannot connect to Postgres", err.Error())
	}

	repo := repository.NewRepos(db)

	services := service.NewServices(repo)

	h := transport.NewHandler(services)

	s := internal.NewServer(*cfg, h.InitRoutes())

	s.Run()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
