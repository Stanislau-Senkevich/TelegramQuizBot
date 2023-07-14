package repository

import (
	"QuizBot/internal/entity"
	"time"
)

type StatisticsRepository interface {
	SaveResults(quiz *entity.Quiz) error
	GetStatistics(chatID int64, timeSinceQuiz time.Duration) ([]entity.StatQuiz, error)
}
