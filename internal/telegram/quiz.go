package telegram

import (
	"QuizBot/internal/entity"
	"QuizBot/internal/keyboard"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sort"
	"time"
)

func (b *Bot) handleStartQuiz(chatID int64) error {
	quiz, err := b.quizRep.GetQuiz(chatID)
	if err != nil {
		return err
	}

	tasks, err := b.taskRep.GetTasksForQuiz(quiz)
	if err != nil {
		return err
	}

	err = b.quizRep.FillQuizWithTasks(chatID, tasks)
	if err != nil {
		return err
	}

	err = b.sendCurrentTask(chatID)
	if err != nil {
		return err
	}

	err = b.userRep.SetStage(chatID, b.Config.QuizStage)
	if err != nil {
		return err
	}
	return err
}

func (b *Bot) handleQuiz(chatID int64) error {
	quiz, err := b.quizRep.GetQuiz(chatID)
	if err != nil {
		return err
	}
	if quiz.CurrentTaskIter >= len(quiz.Tasks) {
		feedBack := fmt.Sprintf(b.Config.QuizResults, quiz.CorrectTasks, quiz.TaskAmount)
		msg := tgbotapi.NewMessage(chatID, feedBack)
		keyboard.MakeKeyBoard(&msg, []string{
			b.Config.QuizButton,
			b.Config.StatisticsButton,
			b.Config.ContactsButton,
			b.Config.AboutButton,
		})
		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}

		err = b.userRep.SetStage(chatID, b.Config.StartStage)
		if err != nil {
			return err
		}

		err = b.statsRep.SaveResults(quiz)
		if err != nil {
			return err
		}

		_, err = b.quizRep.DeleteQuiz(chatID)
		if err != nil {
			return err
		}
		return nil
	}

	err = b.sendCurrentTask(chatID)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) sendCurrentTask(chatID int64) error {
	task, iter, err := b.quizRep.GetCurrentTask(chatID)
	if err != nil {
		return err
	}
	poll := makePoll(task, chatID, iter+1)
	msg, _ := b.bot.Send(poll)
	err = b.logPollInDB(&msg)
	return err
}

func (b *Bot) sendStatistics(chatID int64, timeValue time.Duration) error {
	stats, err := b.statsRep.GetStatistics(chatID, timeValue)
	if err != nil {
		return err
	}

	if len(stats) == 0 {
		msg := tgbotapi.NewMessage(chatID, b.Config.NoStatistics)
		b.makeStartKeyBoard(&msg)
		_, _ = b.bot.Send(msg)
		return nil
	}

	statistics := entity.InitStatistics()
	for _, v := range stats {
		statistics.Spheres[v.Sphere]++
		statistics.Sections[v.Section]++
		statistics.Difficulties[v.Difficulty]++
		statistics.TasksAmount += v.TaskAmount
		statistics.CorrectTasksAmount += v.CorrectTasks
		statistics.QuizAmount++
	}
	msg := tgbotapi.NewMessage(chatID, statistics.Print(b.Config))
	b.makeStartKeyBoard(&msg)
	_, _ = b.bot.Send(msg)
	return nil
}

func isCorrectAnswer(userAns, correctAns []int) bool {
	if len(userAns) != len(correctAns) {
		return false
	}
	sort.Slice(userAns, func(i, j int) bool { return userAns[i] < userAns[j] })
	sort.Slice(correctAns, func(i, j int) bool { return correctAns[i] < correctAns[j] })
	for i := range correctAns {
		if userAns[i] != correctAns[i] {
			return false
		}
	}
	return true
}
