package repository

import "QuizBot/internal/entity"

type PollRepository interface {
	SavePoll(poll *entity.Poll) (bool, error)
	GetPoll(pollID string) (*entity.Poll, error)
	DeletePoll(poll *entity.Poll) (bool, error)
}
