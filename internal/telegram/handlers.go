package telegram

import (
	"context"
	"log/slog"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (b *Bot) handleStart(
	ctx context.Context,
	client *tgbot.Bot,
	update *models.Update,
) {
	if update.Message == nil {
		return
	}

	_, err := client.SendMessage(
		ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Привет! Я бот на Go.",
		})
	if err != nil {
		slog.Error("failed to send start message", "error", err)
	}
}

func (b *Bot) handleHelp(
	ctx context.Context,
	client *tgbot.Bot,
	update *models.Update,
) {
	if update.Message == nil {
		return
	}

	_, err := client.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Доступные команды:\n/start — запуск\n/help — помощь",
	})
	if err != nil {
		slog.Error("failed to send help message", "error", err)
	}
}
