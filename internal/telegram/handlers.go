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

func (b *Bot) handleProfile(
	ctx context.Context,
	client *tgbot.Bot,
	update *models.Update,
) {
	chatID := update.Message.Chat.ID
	tgID := update.Message.From.ID

	// Идём в панель за данными
	item, err := b.panel.ClientByTgId(ctx, tgID)
	if err != nil {
		slog.Error("получение клиента", "tg_id", tgID, "error", err)

		text := "Не удалось получить данные, попробуйте позже."
		if errors.Is(err, xui.ErrClientNotFound) {
			text = "Подписка не найдена 🤷"
		}
		b.api.SendMessage(ctx, &tgbot.SendMessageParams{ChatID: chatID, Text: text})

		b.api.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID:    chatID,
			Text:      "Тестовый текст",
			ParseMode: models.ParseModeHTML,
		})
}
