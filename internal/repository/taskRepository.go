package repository

import (
	"QuizBot/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type TaskRepository interface {
	GetTasksForQuiz(quiz *entity.Quiz) ([]entity.Task, error)
	GetAllRequiredTasks(filter bson.M) ([]entity.Task, error)
	GetAllSpheres() ([]string, error)
	GetAllSectionsOfSphere(sphere string) ([]string, error)
	GetAllDifficultiesOfSelected(sphere string, section string) ([]string, error)
	IsValidSphere(req string) (bool, error)
	IsValidSection(req, sphere string) (bool, error)
	IsValidDifficulty(req, sphere, section string) (bool, error)
	IsValid(req, validKey string, quiz *entity.Quiz) (bool, error)
}
