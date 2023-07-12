package entity

import "time"

type Quiz struct {
	ChatID          int64    `bson:"chat_id"`
	CurrentTaskIter int      `bson:"current_task_iter"`
	TaskAmount      int      `bson:"tasks_amount"`
	CorrectTasks    int      `bson:"correct_tasks"`
	Tasks           []Task   `bson:"tasks"`
	Sphere          string   `bson:"sphere"`
	Section         string   `bson:"section"`
	Tags            []string `bson:"tags"`
	Difficulty      string   `bson:"difficulty"`
}

type StatQuiz struct {
	ChatID       int64     `bson:"chat_id"`
	Sphere       string    `bson:"sphere"`
	Section      string    `bson:"section"`
	Difficulty   string    `bson:"difficulty"`
	TaskAmount   int       `bson:"tasks_amount"`
	CorrectTasks int       `bson:"correct_tasks"`
	Date         time.Time `bson:"date"`
}

func (q *Quiz) GetStatVersion() *StatQuiz {
	return &StatQuiz{
		ChatID:       q.ChatID,
		Sphere:       q.Sphere,
		Section:      q.Section,
		Difficulty:   q.Difficulty,
		TaskAmount:   q.TaskAmount,
		CorrectTasks: q.CorrectTasks,
		Date:         time.Now(),
	}
}
