package config

import (
	"errors"
	"os"
)

// Config содержит настройки приложения
type Config struct {
	TelegramBotToken string
	XRayPanelToken   string
	XRayPanelLink    string
}

// Load читает конфигурацию из переменных окружения
func Load() (Config, error) {
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
