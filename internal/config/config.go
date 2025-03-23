package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env     string `yaml:"env" env-default:"development"`
	Db      `yaml:"db"`
	HttpSrv `yaml:"http"`
	S3      `yaml:"s3"`
}

type HttpSrv struct {
	Address     string
	Timeout     time.Duration
	IdleTimeout time.Duration
}

type Db struct {
	User     string
	Password string
	Name     string
	Port     string
	Host     string
	Sslmode  string
}

type S3 struct {
	KeyID       string `yaml:"key_id"`
	AppKey      string `yaml:"app_key"`
	AuthToken   string `yaml:"auth_token"`
	DownloadUrl string `yaml:"download_url"`
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening file: %s", err)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config: %s", err)
	}

	return &cfg
}
