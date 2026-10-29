package telegram

import (
	"context"
	"fmt"

	tgbot "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type messageSender interface {
	SendMessage(
		ctx context.Context,
		params *tgbot.SendMessageParams,
	) (*models.Message, error)
}

type Bot struct {
	client *tgbot.Bot
	sender messageSender
}

func New(token string) (*Bot, error) {
	client, err := tgbot.New(
		token,
		tgbot.WithDefaultHandler(defaultHandler),
	)
	if err != nil {
		return nil, fmt.Errorf("create telegram bot: %w", err)
	}
	b := &Bot{
		client: client,
		sender: client,
	}
	b.registerHandlers()

	return b, nil
}

func (b *Bot) registerHandlers() {
	b.client.RegisterHandler(
		tgbot.HandlerTypeMessageText,
		"/start",
		tgbot.MatchTypeExact,
		b.handleStart,
	)
	b.client.RegisterHandler(
		tgbot.HandlerTypeMessageText,
		"/help",
		tgbot.MatchTypeExact,
		b.handleHelp,
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
