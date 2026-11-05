package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config содержит настройки приложения
type Config struct {
	TelegramBotToken string
	XRayPanelToken   string
	XRayPanelLink    string
}

// Load читает конфигурацию из переменных окружения
func Load() (Config, error) {
	// Загружаем .env, если файл есть. Если нет — не падаем,
	// возможно переменные заданы в окружении (Docker, systemd и т.д.)
	if err := godotenv.Load(); err != nil {
		log.Print(".env не найден, используются переменные окружения")
	}

	cfg := Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		XRayPanelToken:   os.Getenv("XRAY_TOKEN"),
		XRayPanelLink:    os.Getenv("XRAY_LINK"),
	}
	if cfg.TelegramBotToken == "" {
		return Config{}, errors.New("TELEGRAM_BOT_TOKEN is not set")
	}
	if cfg.XRayPanelToken == "" {
		return Config{}, errors.New("XRAY_TOKEN is not set")
	}
	if cfg.XRayPanelLink == "" {
		return Config{}, errors.New("XRAY_LINK is not set")
	}
	return cfg, nil
}
