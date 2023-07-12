package main

import (
	"QuizBot/pkg/config"
	mongoDB "QuizBot/pkg/repository/mongo"
	"QuizBot/pkg/repository/postgres"
	"QuizBot/pkg/telegram"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	api, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	api.Debug = true

	mongo, err := mongoDB.InitMongoRepository(&cfg.Mongo)
	defer func() {
		if err = mongo.DB.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	postgres, err := postgresDB.InitPostgresRepository(&cfg.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = postgres.DB.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	bot := telegram.NewBot(api, postgres, postgres, mongo, mongo, mongo, cfg)

	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}

}
