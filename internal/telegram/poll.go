package telegram

import (
	"QuizBot/internal/entity"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) stopPoll(update *tgbotapi.Update) (bool, error) {
	poll, err := b.pollRep.GetPoll(update.PollAnswer.PollID)
	if err != nil {
		return false, err
	}
	if poll == nil {
		return false, nil
	}
	_, err = b.bot.StopPoll(tgbotapi.StopPollConfig{
		BaseEdit: tgbotapi.BaseEdit{
			ChatID:          poll.ChatID,
			MessageID:       poll.MessageID,
			ChannelUsername: "",
			InlineMessageID: "",
			ReplyMarkup:     nil,
		}},
	)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *Bot) logPollInDB(msg *tgbotapi.Message) error {
	poll := entity.Poll{PollID: msg.Poll.ID, ChatID: msg.Chat.ID, MessageID: msg.MessageID}
	_, err := b.pollRep.SavePoll(&poll)
	return err
}

func (b *Bot) sendFeedBack(result bool, chatID int64, task *entity.Task) error {
	if result {
		_, _ = b.bot.Send(tgbotapi.NewMessage(chatID, b.Config.QuizCorrect))
	} else {
		correct := "\nCorrect answer:\n"
		for _, v := range task.CorrectOptionIDs {
			correct += fmt.Sprintf("%d) %s", v+1, task.Options[v])
		}
		_, _ = b.bot.Send(tgbotapi.NewMessage(chatID, b.Config.QuizWrong+correct))
	}
	return nil
}

func (b *Bot) handlePollAnswer(update *tgbotapi.Update) error {
	poll, err := b.pollRep.GetPoll(update.PollAnswer.PollID)
	if err != nil {
		return err
	}

	task, _, err := b.quizRep.GetCurrentTask(poll.ChatID)
	if err != nil {
		return err
	}

	result := isCorrectAnswer(update.PollAnswer.OptionIDs, task.CorrectOptionIDs)

	err = b.sendFeedBack(result, poll.ChatID, task)
	if err != nil {
		return err
	}
	err = b.quizRep.UpdateQuiz(poll.ChatID, result)
	if err != nil {
		return err
	}

	err = b.handleQuiz(poll.ChatID)
	if err != nil {
		return err
	}

	_, err = b.stopPoll(update)
	if err != nil {
		return err
	}
	_, err = b.pollRep.DeletePoll(poll)
	if err != nil {
		return err
	}
	return nil
}

func makePoll(task *entity.Task, chatID int64, prefixNumber int) *tgbotapi.SendPollConfig {
	prefix := fmt.Sprintf("Question #%d\n", prefixNumber)
	poll := tgbotapi.NewPoll(chatID, prefix+task.Question, task.Options...)
	poll.AllowsMultipleAnswers = true
	poll.IsAnonymous = false
	return &poll
}
