package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleUpdates(updates *tgbotapi.UpdatesChannel) {
	for update := range *updates {
		b.handleUpdate(&update)
	}
}

func (b *Bot) handleUpdate(update *tgbotapi.Update) {
	if update.Message != nil && update.Message.Text != "" { //nolint
		chatID := update.Message.Chat.ID
		user, err := b.userRep.CheckUser(chatID)
		if err != nil {
			b.handleError(update, err)
		}
		if !user {
			err = b.userRep.LogUser(chatID)
			if err != nil {
				b.handleError(update, err)
			}
		}

		if b.isCommand(update.Message.Text) {
			b.handleCommand(update)
		} else {
			b.handleMessage(update)
		}
	} else if update.PollAnswer != nil {
		err := b.handlePollAnswer(update)
		if err != nil {
			b.handleError(update, err)
		}
	}
}

func (b *Bot) handleMessage(update *tgbotapi.Update) { //nolint
	stage, err := b.userRep.GetStage(update.Message.Chat.ID)
	if err != nil {
		b.handleError(update, err)
	}
	switch stage {
	case b.Config.SphereStage:
		err = b.handleSettings(update, b.Config.Mongo.Quiz.SphereField, b.Config.SectionStage)
		if err != nil {
			b.handleError(update, err)
		}
	case b.Config.SectionStage:
		err = b.handleSettings(update, b.Config.Mongo.Quiz.SectionField, b.Config.DifficultyStage)
		if err != nil {
			b.handleError(update, err)
		}
	case b.Config.DifficultyStage:
		err = b.handleSettings(update, b.Config.Mongo.Quiz.DifficultyField, b.Config.AmountStage)
		if err != nil {
			b.handleError(update, err)
		}
	case b.Config.AmountStage:
		err = b.handleSettings(update, b.Config.Mongo.Quiz.TaskAmountField, b.Config.QuizStage)
		if err != nil {
			b.handleError(update, err)
			return
		}
		err = b.handleStartQuiz(update.Message.Chat.ID)
		if err != nil {
			b.handleError(update, err)
		}
	case b.Config.StatisticsStage:
		duration, err := b.getDuration(update.Message.Text)
		if err != nil {
			b.handleError(update, err)
		}
		err = b.sendStatistics(update.Message.Chat.ID, duration)
		if err != nil {
			b.handleError(update, err)
		}
		err = b.userRep.SetStage(update.Message.Chat.ID, b.Config.StartStage)
		if err != nil {
			b.handleError(update, err)
		}
	}
}

func (b *Bot) handleCommand(update *tgbotapi.Update) { //nolint
	chatID := update.Message.Chat.ID

	if update.Message.Text == b.Config.StopButton {
		err := b.handleStop(chatID)
		if err != nil {
			b.handleError(update, err)
		}
		return
	}

	isStart, err := b.stageBeforeCommand(chatID)
	if err != nil {
		b.handleError(update, err)
	}
	if !isStart {
		return
	}

	switch update.Message.Text {
	case b.Config.StartCommand:
		err = b.handleStart(chatID)
		if err != nil {
			b.handleError(update, err)
		}
	case b.Config.QuizButton:
		err = b.quizRep.InitEmptyQuiz(chatID)
		if err != nil {
			b.handleError(update, err)
		}
		err = b.handleSettings(update, b.Config.Mongo.Quiz.NoneValue, b.Config.SphereStage)
		if err != nil {
			b.handleError(update, err)
		}
	case b.Config.StatisticsButton:
		err = b.handleStatistics(chatID)
		if err != nil {
			b.handleError(update, err)
		}
	case b.Config.AboutButton:
		_, _ = b.bot.Send(tgbotapi.NewMessage(chatID, b.Config.AboutReply))
	case b.Config.ContactsButton:
		_, _ = b.bot.Send(tgbotapi.NewMessage(chatID, b.Config.ContactsReply))
	}
}

func (b *Bot) stageBeforeCommand(chatID int64) (bool, error) {
	stage, err := b.userRep.GetStage(chatID)
	if err != nil {
		return false, err
	}
	if stage != b.Config.StartStage {
		_, err = b.bot.Send(tgbotapi.NewMessage(chatID, b.Config.StopQuiz))
		if err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
