package telegram

import (
	"3xuibot/internal/xui"
	"context"
	"fmt"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Bot представляет Telegram-бота нашего приложения
type Bot struct {
	api   *tgbot.Bot
	panel *xui.ClientAPI
}

// New создаёт Telegram-бота и регистрирует обработчики
func New(token string, panel *xui.ClientAPI) (*Bot, error) {
	api, err := tgbot.New(
		token,
		tgbot.WithDefaultHandler(defaultHandler),
	)
	if err != nil {
		return nil, fmt.Errorf("create telegram bot: %w", err)
	}
	b := &Bot{
		api:   api,
		panel: panel,
	}
	b.registerHandlers()

	return b, nil
}

// Start запускает получение обновлений от Telegram
func (b *Bot) Start(ctx context.Context) {
	b.api.Start(ctx)
}

// Регистрируем все обработчики
func (b *Bot) registerHandlers() {
	b.api.RegisterHandler(
		tgbot.HandlerTypeMessageText,
		"/start",
		tgbot.MatchTypeExact,
		b.handleStart,
	)
	b.api.RegisterHandler(
		tgbot.HandlerTypeMessageText,
		"/help",
		tgbot.MatchTypeExact,
		b.handleHelp,
	)
	b.api.RegisterHandler(
		tgbot.HandlerTypeMessageText,
		"/profile",
		tgbot.MatchTypeExact,
		b.handleProfile,
	)

	// reply-кнопки (приходят как обычный текст)
	// Кнопка получния профиля
	b.api.RegisterHandler(
		tgbot.HandlerTypeMessageText,
		btnProfile,
		tgbot.MatchTypeExact,
		b.handleProfile,
	)
	// Кнопка создания подписи
	b.api.RegisterHandler(
		tgbot.HandlerTypeMessageText,
		btnCreate,
		tgbot.MatchTypeExact,
		b.handleCreateSubscription,
	)
}

func defaultHandler(
	ctx context.Context,
	client *tgbot.Bot,
	update *models.Update,
) {
	if update.Message == nil || update.Message.Text == "" {
		return
	}

	_, err := client.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Ты написал" + update.Message.Text,
	})
	if err != nil {
		fmt.Printf("send message: %v\n", err)
	}
}
