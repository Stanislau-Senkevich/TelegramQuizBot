package telegram

import (
	"QuizBot/internal/config"
	"QuizBot/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	userRep  repository.UserRepository
	pollRep  repository.PollRepository
	taskRep  repository.TaskRepository
	quizRep  repository.QuizRepository
	statsRep repository.StatisticsRepository

	Config *config.Config
}

func NewBot(bot *tgbotapi.BotAPI,
	userRep repository.UserRepository,
	pollRep repository.PollRepository,
	taskRep repository.TaskRepository,
	quizRep repository.QuizRepository,
	statsRep repository.StatisticsRepository,
	cfg *config.Config,
) *Bot {
	return &Bot{bot: bot,
		userRep:  userRep,
		pollRep:  pollRep,
		taskRep:  taskRep,
		quizRep:  quizRep,
		statsRep: statsRep,
		Config:   cfg,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	updates := b.initUpdatesChannel()
	b.handleUpdates(&updates)
	return nil
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}
