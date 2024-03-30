package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string `env:"APP_NAME" envDefault:"CLI_MSG_Service"`

	AppVersion string // `env:"APP_VER"`
	URL        string `env:"URL"`
	Resp       string `env:"RESP"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: - %v ", err)
	}

	cfg.AppName = os.Getenv("APP_NAME")
	cfg.AppVersion = "0.0.1" // os.Getenv("APP_VER")
	cfg.URL = os.Getenv("URL")
	cfg.Resp = os.Getenv("RESP")

	return &cfg, nil
}
