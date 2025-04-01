package main

import (
	"fmt"
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
	"os/signal"
	"syscall"
)

// @title           Kolesa API
// @version         1.0
// @description     This is a sample server celler server.
// @host      kolesaapi.onrender.com
// @BasePath  /api
func main() {
	cfg := config.LoadConfig()

	log := setupLogger(cfg.Env)
	log.Info("Start the application, config loaded")

	db, err := postgres.NewPostgresDB(cfg.Db.User, cfg.Db.Name, cfg.Db.Port, cfg.Db.Password, cfg.Db.Host, cfg.Db.Sslmode)
	if err != nil {
		log.Error("Cannot connect to Postgres", err.Error())
	}
	log.Info("Postgres started:", "HOST ADDRESS", (fmt.Sprintf("%s:%s", cfg.Db.Host, cfg.Db.Port)))

	s3, err := pkg.NewS3Client(cfg.S3.KeyID, cfg.S3.BucketID, cfg.S3.AppKey, cfg.S3.AuthToken, cfg.S3.ApiUrl, cfg.S3.DownloadUrl, cfg.S3.UploadUrl, log)
	if err != nil {
		log.Error("Cannot connect to S3", err.Error())
	}
	log.Info("S3 started with", slog.String("API_URL", cfg.S3.ApiUrl))

	repo := repository.NewRepos(db)

	services := service.NewServices(repo, s3)

	h := transport.NewHandler(services, repo, s3)

	s := internal.NewServer(*cfg, h.InitRoutes())

	go func() {
		err := s.Run()
		if err != nil {
			log.Error("Cannot start HTTP server", err.Error())
		}
	}()

	log.Info("Server is running", "ADDRESS:", cfg.HttpSrv.Address)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
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
