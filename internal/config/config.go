package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env     string `env:"ENV" env-default:"development"`
	Db      Db
	HttpSrv HttpSrv
	GRPC    GRPC
	S3      S3
	JWT     JWT
}

type HttpSrv struct {
	Address     string        `env:"HTTP_ADDRESS"`
	Timeout     time.Duration `env:"HTTP_TIMEOUT"`
	IdleTimeout time.Duration `env:"HTTP_IDLE_TIMEOUT"`
}

type GRPC struct {
	Address      string        `env:"GRPC_ADDRESS"`
	Timeout      time.Duration `env:"GRPC_TIMEOUT"`
	RetriesCount int           `env:"GRPC_RETRIES_COUNT"`
}

type JWT struct {
	JWTSecret string `env:"JWT_SECRET"`
}

type Db struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
	Host     string `env:"DB_HOST"`
	Sslmode  string `env:"DB_SSLMODE"`
}

type S3 struct {
	KeyID       string `env:"S3_KEY_ID"`
	BucketID    string `env:"S3_BUCKET_ID"`
	AppKey      string `env:"S3_APP_KEY"`
	AuthToken   string `env:"S3_AUTH_TOKEN"`
	ApiUrl      string `env:"S3_API_URL"`
	DownloadUrl string `env:"S3_DOWNLOAD_URL"`
	UploadUrl   string `env:"S3_UPLOAD_URL"`
}

func LoadConfig() *Config {
	// Not used in production, only development
	godotenv.Load()

	var cfg Config

	cfg.Env = os.Getenv("ENV")

	cfg.Db.User = os.Getenv("DB_USER")
	cfg.Db.Password = os.Getenv("DB_PASSWORD")
	cfg.Db.Name = os.Getenv("DB_NAME")
	cfg.Db.Port = os.Getenv("DB_PORT")
	cfg.Db.Host = os.Getenv("DB_HOST")
	cfg.Db.Sslmode = os.Getenv("DB_SSLMODE")

	cfg.HttpSrv.Address = os.Getenv("HTTP_ADDRESS")
	cfg.HttpSrv.Timeout, _ = time.ParseDuration(os.Getenv("HTTP_TIMEOUT"))
	cfg.HttpSrv.IdleTimeout, _ = time.ParseDuration(os.Getenv("HTTP_IDLE_TIMEOUT"))

	cfg.S3.KeyID = os.Getenv("S3_KEY_ID")
	cfg.S3.BucketID = os.Getenv("S3_BUCKET_ID")
	cfg.S3.AppKey = os.Getenv("S3_APP_KEY")
	cfg.S3.AuthToken = os.Getenv("S3_AUTH_TOKEN")
	cfg.S3.ApiUrl = os.Getenv("S3_API_URL")
	cfg.S3.DownloadUrl = os.Getenv("S3_DOWNLOAD_URL_URL")
	cfg.S3.UploadUrl = os.Getenv("S3_UPLOAD_URL")

	cfg.GRPC.Address = os.Getenv("GRPC_ADDRESS")
	cfg.GRPC.Timeout, _ = time.ParseDuration(os.Getenv("GRPC_TIMEOUT"))
	cfg.GRPC.RetriesCount, _ = strconv.Atoi(os.Getenv("GRPC_RETRIES_COUNT"))

	cfg.JWT.JWTSecret = os.Getenv("JWT_SECRET")

	return &cfg
}
