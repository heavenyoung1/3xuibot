package telegram

import "github.com/go-telegram/bot/models"

const btnMenu = "📋 Меню"
const btnHelp = "❓ Помощь "

// Обычная клавиатура под полем ввода
func mainKeyboard() *models.ReplyKeyboardMarkup {
	return &models.ReplyKeyboardMarkup{
		Keyboard: [][]models.KeyboardButton{
			{
				{Text: btnMenu}, {Text: btnHelp}, // одна строка - две кнопки
			},
		},
		ResizeKeyboard: true,
	}
}

// Inline-кнопки под сообщением
func menuKeyboard() *models.InlineKeyboardMarkup {
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Пункт 1", CallbackData: "item:1"},
				{Text: "Пункт 2", CallbackData: "item:2"},
			},
			{
				{Text: "Открыть сайт", URL: "https://go.dev"},
			},
		},
	}
}
