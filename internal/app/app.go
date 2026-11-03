package app

import (
	"3xuibot/internal/config"
	"3xuibot/internal/telegram"
	"3xuibot/internal/xui"
	"context"
)

// App представляет приложение целиком
type Application struct {
	panel *xui.ClientAPI
	bot   *telegram.Bot
}

// New создаёт приложение и все необходимые зависимости
func New(cfg config.Config) (*Application, error) {
	panel := xui.NewClientAPI(
		cfg.XRayPanelLink,
		cfg.XRayPanelToken,
	)

	bot, err := telegram.New(
		cfg.TelegramBotToken,
		panel,
	)
	if err != nil {
		return nil, err
	}

	return &Application{
		bot:   bot,
		panel: panel,
	}, nil
}

// Run запускает приложение
func (a *Application) Run(ctx context.Context) {
	a.bot.Start(ctx)
}
