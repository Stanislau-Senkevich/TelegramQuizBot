package postgresdb

import (
	"QuizBot/internal/entity"
	botError "QuizBot/internal/error"
	"errors"
	_ "github.com/lib/pq" //nolint
)

func (p *PostgresRepository) SavePoll(poll *entity.Poll) (bool, error) {
	_, err := p.DB.Exec("INSERT INTO polls (poll_id, chat_id, message_id) values ($1, $2, $3);",
		poll.PollID, poll.ChatID, poll.MessageID)
	if err != nil {
		return false, botError.NewBotError(err, p.Config.BotProblem)
	}
	return true, nil
}

func (p *PostgresRepository) GetPoll(pollID string) (*entity.Poll, error) {
	poll := make([]entity.Poll, 0)
	err := p.DB.Select(&poll, "SELECT * FROM polls where poll_id = $1", pollID)
	if err != nil {
		return nil, botError.NewBotError(err, p.Config.BotProblem)
	}
	if len(poll) == 0 {
		return nil, botError.NewBotError(errors.New("nil poll"), p.Config.BotProblem)
	}
	return &poll[0], nil
}

func (p *PostgresRepository) DeletePoll(poll *entity.Poll) (bool, error) {
	res, err := p.DB.Exec("DELETE FROM polls where poll_id = $1", poll.PollID)
	if err != nil {
		return false, botError.NewBotError(err, p.Config.BotProblem)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return false, botError.NewBotError(err, p.Config.BotProblem)
	}
	if rows > 0 {
		return true, nil
	}
	return false, nil
}
