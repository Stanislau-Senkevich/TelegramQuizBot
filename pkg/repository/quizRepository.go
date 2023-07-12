package repository

import "QuizBot/pkg/entity"

type QuizRepository interface {
	AddQuiz(quiz *entity.Quiz) error
	UpdateQuiz(chatID int64, isCorrectAns bool) error
	GetCurrentTask(chatID int64) (*entity.Task, int, error)
	DeleteQuiz(chatID int64) (bool, error)
	GetQuiz(chatID int64) (*entity.Quiz, error)
	InitEmptyQuiz(chatID int64) error
	UpdateQuizSettings(chatID int64, key, value string) error
	SetQuizTasks(chatID int64, tasks []entity.Task) error
	FillQuizWithTasks(chatID int64, tasks []entity.Task) error
}
