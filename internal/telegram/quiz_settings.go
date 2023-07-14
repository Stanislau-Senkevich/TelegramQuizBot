package telegram

import (
	"QuizBot/internal/entity"
	botError "QuizBot/internal/error"
	"QuizBot/internal/keyboard"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"time"
)

func (b *Bot) makeKeyWords(stage string, quiz *entity.Quiz) ([]string, error) {
	var keys []string
	var err error

	switch stage {
	case b.Config.SphereStage:
		keys, err = b.taskRep.GetAllSpheres()
		if err != nil {
			return nil, err
		}
		keys = append(keys, b.Config.AllButton, b.Config.StopButton)
	case b.Config.SectionStage:
		keys, err = b.taskRep.GetAllSectionsOfSphere(quiz.Sphere)
		if err != nil {
			return nil, err
		}
		keys = append(keys, b.Config.AllButton, b.Config.StopButton)
	case b.Config.DifficultyStage:
		keys, err = b.taskRep.GetAllDifficultiesOfSelected(quiz.Sphere, quiz.Section)
		if err != nil {
			return nil, err
		}
		keys = append(keys, b.Config.AllButton, b.Config.StopButton)
	case b.Config.AmountStage:
		quiz.TaskAmount = 101
		tasks, err := b.taskRep.GetTasksForQuiz(quiz)
		if err != nil {
			return nil, err
		}
		keys = getNumbersForAmountButtons(len(tasks))
		keys = append(keys, b.Config.StopButton)
	case b.Config.QuizStage:
		keys = []string{b.Config.StopButton}
	}
	return keys, nil
}

func getNumbersForAmountButtons(length int) []string {
	var keys []string
	if length < 5 {
		keys = append(keys, strconv.Itoa(length))
	}

	if length >= 5 {
		keys = append(keys, "5")
	}

	if length >= 10 {
		keys = append(keys, "10")
	}

	if length >= 15 {
		keys = append(keys, "15")
	}

	if length >= 20 {
		keys = append(keys, "20")
	}

	if length >= 30 {
		keys = append(keys, "30")
	}

	if length >= 40 {
		keys = append(keys, "40")
	}

	if length >= 50 {
		keys = append(keys, "50")
	}

	if length >= 100 {
		keys = append(keys, "100")
	}
	return keys
}

func (b *Bot) handleSettings(update *tgbotapi.Update, typeParam string, newStage string) error {
	chatID := update.Message.Chat.ID
	request := update.Message.Text

	quiz, err := b.quizRep.GetQuiz(chatID)
	if err != nil {
		return err
	}

	isValid, err := b.taskRep.IsValid(request, typeParam, quiz)
	if !isValid {
		if err == nil {
			err = errors.New("invalid option was selected")
		}
		return botError.NewBotError(err, b.Config.InvalidOption)
	}

	err = b.quizRep.UpdateQuizSettings(chatID, typeParam, request)
	if err != nil {
		return err
	}

	quiz, err = b.quizRep.GetQuiz(chatID)
	if err != nil {
		return err
	}

	keys, err := b.makeKeyWords(newStage, quiz)
	if err != nil {
		return err
	}

	msgText, err := b.getSettingMessage(newStage)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(chatID, msgText)
	keyboard.MakeKeyBoard(&msg, keys)
	_, _ = b.bot.Send(msg)

	if newStage != b.Config.QuizStage {
		err = b.userRep.SetStage(chatID, newStage)
	}
	return err
}

func (b *Bot) getSettingMessage(newStage string) (string, error) {
	switch newStage {
	case b.Config.SphereStage:
		return b.Config.SelectSphere, nil
	case b.Config.SectionStage:
		return b.Config.SelectSection, nil
	case b.Config.DifficultyStage:
		return b.Config.SelectDifficulty, nil
	case b.Config.AmountStage:
		return b.Config.SelectAmount, nil
	case b.Config.QuizStage:
		return b.Config.QuizStarted, nil
	default:
		return "", botError.NewBotError(nil, b.Config.UnknownError)
	}
}

func (b *Bot) getDuration(txt string) (time.Duration, error) {
	switch txt {
	case b.Config.Days7:
		return 7 * 24 * time.Hour, nil
	case b.Config.Days14:
		return 14 * 24 * time.Hour, nil
	case b.Config.Days30:
		return 30 * 24 * time.Hour, nil
	case b.Config.Days90:
		return 90 * 24 * time.Hour, nil
	case b.Config.AllTime:
		return 10000 * 24 * time.Hour, nil
	}
	return 0 * time.Hour, botError.NewBotError(errors.New("invalid period"), b.Config.InvalidTime)
}
