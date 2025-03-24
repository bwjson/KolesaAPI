package config

import (
	"os"
	"time"
)

type Config struct {
	Env     string `env:"ENV" env-default:"development"`
	Db      Db
	HttpSrv HttpSrv
	S3      S3
}

type HttpSrv struct {
	Address     string        `env:"HTTP_ADDRESS"`
	Timeout     time.Duration `env:"HTTP_TIMEOUT"`
	IdleTimeout time.Duration `env:"HTTP_IDLE_TIMEOUT"`
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
	AppKey      string `env:"S3_APP_KEY"`
	AuthToken   string `env:"S3_AUTH_TOKEN"`
	DownloadUrl string `env:"S3_DOWNLOAD_URL"`
}

func LoadConfig() *Config {
	// Загружаем .env файл
	//if err := godotenv.Load(); err != nil {
	//	log.Fatal("Error loading .env file")
	//}

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
	cfg.S3.AppKey = os.Getenv("S3_APP_KEY")
	cfg.S3.AuthToken = os.Getenv("S3_AUTH_TOKEN")
	cfg.S3.DownloadUrl = os.Getenv("S3_DOWNLOAD_URL")

	return &cfg
}

//type Config struct {
//	Env     string `yaml:"env" env-default:"development"`
//	Db      `yaml:"db"`
//	HttpSrv `yaml:"http"`
//	S3      `yaml:"s3"`
//}
//
//type HttpSrv struct {
//	Address     string
//	Timeout     time.Duration
//	IdleTimeout time.Duration
//}
//
//type Db struct {
//	User     string
//	Password string
//	Name     string
//	Port     string
//	Host     string
//	Sslmode  string
//}
//
//type S3 struct {
//	KeyID       string `yaml:"key_id"`
//	AppKey      string `yaml:"app_key"`
//	AuthToken   string `yaml:"auth_token"`
//	DownloadUrl string `yaml:"download_url"`
//}
//
//func LoadConfig() *Config {
//	if err := godotenv.Load(); err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	configPath := os.Getenv("CONFIG_PATH")
//
//	if configPath == "" {
//		log.Fatal("CONFIG_PATH environment variable not set")
//	}
//
//	if _, err := os.Stat(configPath); err != nil {
//		log.Fatalf("error opening file: %s", err)
//	}
//
//	var cfg Config
//
//	err := cleanenv.ReadConfig(configPath, &cfg)
//	if err != nil {
//		log.Fatalf("error reading config: %s", err)
//	}
//
//	return &cfg
//}
