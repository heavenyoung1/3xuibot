package app

import (
	"3xuibot/internal/config"
	"3xuibot/internal/telegram"

	"example.com/mybot/internal/config"
	"example.com/mybot/internal/telegram"
)

// App представляет приложение целиком
type App struct {
	bot *telegram.Bot
}

// New создаёт приложение и все необходимые зависимости
func New(cfg config.Config) (*App, error) {
	bot, err := telegram.New(cfg.TelegramBotToken)
	if err != nil {
		return nil, err
	}
	return &App{
		bot: bot,
	}, nil
}
