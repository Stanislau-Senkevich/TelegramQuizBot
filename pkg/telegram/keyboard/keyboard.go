package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func makeRowsForKeyBoard(keys []string) [][]tgbotapi.KeyboardButton {
	i := 0
	rows := make([][]tgbotapi.KeyboardButton, 0)

	for i+1 < len(keys) {
		rows = append(rows, []tgbotapi.KeyboardButton{tgbotapi.NewKeyboardButton(keys[i]), tgbotapi.NewKeyboardButton(keys[i+1])})
		i += 2
	}

	if i < len(keys) {
		rows = append(rows, []tgbotapi.KeyboardButton{tgbotapi.NewKeyboardButton(keys[i])})
	}
	return rows
}

func MakeKeyBoard(message *tgbotapi.MessageConfig, keywords []string) {
	rows := makeRowsForKeyBoard(keywords)
	keyboard := tgbotapi.NewReplyKeyboard(rows...)
	message.ReplyMarkup = &keyboard
}
