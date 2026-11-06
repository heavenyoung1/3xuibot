package telegram

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

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
	userName := update.Message.From.Username

	// Получаем доступные InboundIDs
	// Проверяем есть ли у пользователя уже подписка
	isExist, err := b.panel.ClientByTgId(ctx, tgID)
	if err != nil && !errors.Is(err, xui.ErrClientNotFound) {
		slog.Error("получение клиента", "tgID", tgID, "error", err)

		text := "Не удалось получить данные, попробуйте позже."
		client.SendMessage(ctx, &tgbot.SendMessageParams{ChatID: chatID, Text: text})
		return
	}
	slog.Info("клиент получен", "isExist", isExist)
	if err != nil && isExist != nil {
		client.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: chatID,
			Text:   "Подписка уже есть",
		})
		return
	}

	// Вычисление времени
	expiry := time.Now().AddDate(0, 0, 31)
	expiryMs := expiry.UnixMilli()

	// Собираем данные для запроса
	internalData := xui.ClientAddInternal{
		Email:      userName,
		TotalGB:    0,
		ExpiryTime: expiryMs,
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

	text := fmt.Sprintf(
		"🎉 <b>Подписка успешно создана!</b>\n\n"+
			"👤 <b>Пользователь:</b> <code>%s</code>\n"+
			"📦 <b>Трафик:</b> ♾ Безлимит\n"+
			"📅 <b>Действует до:</b> %s\n"+
			"🌐 <b>Лимит IP:</b> %d\n\n"+
			"🔗 <i>Данные для подключения доступны в личном кабинете.</i>",
		userName,
		expiry.Format("02.01.2006 15:04"),
		5,
	)
	client.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID:    chatID,
		Text:      text,
		ParseMode: "HTML",
	})
}
