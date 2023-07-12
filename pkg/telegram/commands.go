package telegram

import (
	"QuizBot/pkg/telegram/keyboard"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) isCommand(msgText string) bool {
	commands := []string{
		b.Config.QuizButton,
		b.Config.StatisticsButton,
		b.Config.ContactsButton,
		b.Config.AboutButton,
		b.Config.StopButton,
		b.Config.StartCommand,
	}
	for _, s := range commands {
		if s == msgText {
			return true
		}
	}
	return false
}

func (b *Bot) handleStart(chatID int64) error {
	err := b.userRep.LogUser(chatID)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(chatID, b.Config.StartReply)
	b.makeStartKeyBoard(&msg)
	_, _ = b.bot.Send(msg)
	return nil
}

func (b *Bot) handleStop(chatID int64) error {
	stage, err := b.userRep.GetStage(chatID)
	if err != nil {
		return err
	}

	var reply string
	if stage != b.Config.StartStage {
		_, err = b.quizRep.DeleteQuiz(chatID)
		if err != nil {
			return err
		}
		err = b.userRep.SetStage(chatID, b.Config.StartStage)
		if err != nil {
			return err
		}
		reply = b.Config.QuizStopped
	} else {
		reply = b.Config.NoQuizToStop
	}
	msg := tgbotapi.NewMessage(chatID, reply)
	keyboard.MakeKeyBoard(&msg, []string{
		b.Config.QuizButton,
		b.Config.StatisticsButton,
		b.Config.ContactsButton,
		b.Config.AboutButton,
	})
	_, _ = b.bot.Send(msg)
	return nil
}

func (b *Bot) handleStatistics(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, b.Config.DurationReply)
	keyboard.MakeKeyBoard(&msg, []string{
		b.Config.Days7,
		b.Config.Days14,
		b.Config.Days30,
		b.Config.Days90,
		b.Config.AllTime,
	})
	_, _ = b.bot.Send(msg)
	err := b.userRep.SetStage(chatID, b.Config.StatisticsStage)
	return err
}

func (b *Bot) makeStartKeyBoard(msg *tgbotapi.MessageConfig) {
	keyboard.MakeKeyBoard(msg, []string{
		b.Config.QuizButton,
		b.Config.StatisticsButton,
		b.Config.ContactsButton,
		b.Config.AboutButton,
	})
}
