package main

import (
	_ "github.com/bwjson/kolesa_api/docs"
	"github.com/bwjson/kolesa_api/internal"
	"github.com/bwjson/kolesa_api/internal/config"
	"github.com/bwjson/kolesa_api/internal/postgres"
	"github.com/bwjson/kolesa_api/internal/repository"
	"github.com/bwjson/kolesa_api/internal/service"
	"github.com/bwjson/kolesa_api/internal/transport"
	"github.com/bwjson/kolesa_api/pkg"
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
	log.Info("Start the application", slog.String("address", cfg.HttpSrv.Address))

	db, err := postgres.NewPostgresDB(cfg.Db.User, cfg.Db.Name, cfg.Db.Port, cfg.Db.Password, cfg.Db.Host, cfg.Db.Sslmode)
	if err != nil {
		log.Error("Cannot connect to Postgres", err.Error())
	}

	s3, err := pkg.NewS3Client(cfg.S3.KeyID, cfg.S3.AppKey, cfg.S3.AuthToken, cfg.S3.DownloadUrl, log)
	if err != nil {
		log.Error("Cannot connect to S3", err.Error())
	}

	log.Info("S3", cfg)

	repo := repository.NewRepos(db)

	services := service.NewServices(repo)

	h := transport.NewHandler(services, s3)

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
