package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (b *Bot) handleError(update *tgbotapi.Update, err error) {
	var chatID int64
	if update.Message != nil {
		chatID = update.Message.Chat.ID
	} else if update.PollAnswer != nil {
		poll, err := b.pollRep.GetPoll(update.PollAnswer.PollID)
		if err != nil {
			return
		}
		chatID = poll.ChatID
	}
	split := strings.SplitAfter(err.Error(), "\n")
	_, _ = b.bot.Send(tgbotapi.NewMessage(chatID, split[len(split)-1]))
}
