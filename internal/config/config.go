package config

import (
	"os"
	"strings"
)

// Config содержит настройки приложения
type Config struct {
	TelegramBotToken string
	XRayPanelTokens  []string
	XRayPanelLinks   []string
}

// Load читает конфигурацию из переменных окружения
func Load() (Config, error) {
	cfg := Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		XRayPanelTokens:  strings.Split(os.Getenv("XRAY_TOKENS"), ","),
		XRayPanelLinks:   strings.Split(os.Getenv("XRAY_LINKS"), ","),
	}
	return cfg, nil
}
