package telegram

import (
	"context"
	"errors"
	"log/slog"

	"3xuibot/internal/xui"

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
			ChatID:      update.Message.Chat.ID,
			Text:        "Привет! Я бот на Go.",
			ReplyMarkup: mainKeyboard(),
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
		slog.Error("получение клиента", "tgID", tgID, "error", err)

		text := "Не удалось получить данные, попробуйте позже."
		if errors.Is(err, xui.ErrClientNotFound) {
			text = "Подписка не найдена 🤷"
		}
		client.SendMessage(ctx, &tgbot.SendMessageParams{ChatID: chatID, Text: text})
		return
	}

	client.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID:    chatID,
		Text:      "Тестовый текст" + item.Client.Email,
		ParseMode: models.ParseModeHTML,
	})
}

func (b *Bot) handleCreateSubscription(
	ctx context.Context,
	client *tgbot.Bot,
	update *models.Update,
) {
	chatID := update.Message.Chat.ID
	tgID := update.Message.From.ID

	// Получаем доступные InboundIDs

	// Собираем данные для запроса
	internalData := xui.ClientAddInternal{
		Email:      "string",
		TotalGB:    500,
		ExpiryTime: 3000000,
		TgID:       int32(tgID),
		LimitIP:    5,
		Enable:     true,
	}
	dataRequest := xui.ClientAddRequest{
		Client:     internalData,
		InboundIDs: []int{1, 2, 3},
	}
	// Идём в панель
	res, err := b.panel.AddClient(ctx, dataRequest)
	if err != nil {
		slog.Error("создание подписки", "tgID", tgID, "error", err)

		text := "Не удалось создать подписку, попробуйте позже."
		if errors.Is(err, xui.ErrClientNotAdded) {
			text = "Подписка не создана 🤷"
		}
		client.SendMessage(ctx, &tgbot.SendMessageParams{ChatID: chatID, Text: text})
		return
	}
	if !res {
		slog.Warn("AddClient вернул false без ошибки", "tgID", tgID)
		client.SendMessage(ctx, &tgbot.SendMessageParams{ChatID: chatID, Text: "Подписка не создана 🤷"})
		return
	}

	client.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: chatID,
		Text:   "подписка создана",
	})
}
