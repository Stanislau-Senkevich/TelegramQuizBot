package postgresDB

import (
	"QuizBot/pkg/entity"
	botError "QuizBot/pkg/error"
	"errors"
)

func (p *PostgresRepository) LogUser(chatID int64) error {
	c, err := p.CheckUser(chatID)
	if err != nil {
		return err
	}

	if !c {
		_, err = p.DB.Exec("INSERT INTO users (chat_id, stage) values ($1, $2)", chatID, p.Config.StartStage)
	}

	if err != nil {
		return botError.NewBotError(err, p.Config.BotProblem)
	}
	return nil
}

func (p *PostgresRepository) CheckUser(chatID int64) (bool, error) {
	var users []entity.User

	err := p.DB.Select(&users, "SELECT * FROM users where chat_id = $1", chatID)
	if err != nil {
		return false, botError.NewBotError(err, p.Config.BotProblem)
	}

	return len(users) > 0, nil
}

func (p *PostgresRepository) GetStage(chatID int64) (string, error) {
	var users []entity.User

	err := p.DB.Select(&users, "SELECT * FROM users where chat_id = $1", chatID)
	if err != nil {
		return "", botError.NewBotError(err, p.Config.BotProblem)
	}

	if len(users) == 0 {
		return "", botError.NewBotError(errors.New("user wasn't found"), p.Config.BotProblem)
	}

	return users[0].Stage, nil
}

func (p *PostgresRepository) SetStage(chatID int64, stage string) error {
	c, err := p.CheckUser(chatID)
	if err != nil {
		return err
	}

	if !c {
		_, err = p.DB.Exec("INSERT INTO users (chat_id, stage) values ($1, $2)", chatID, stage)
		if err != nil {
			return botError.NewBotError(err, p.Config.BotProblem)
		}
	} else {
		_, err = p.DB.Exec("UPDATE users SET stage = $2 WHERE chat_id = $1", chatID, stage)
		if err != nil {
			return botError.NewBotError(err, p.Config.BotProblem)
		}
	}
	return nil
}
