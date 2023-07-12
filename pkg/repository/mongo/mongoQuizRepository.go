package mongoDB

import (
	"QuizBot/pkg/entity"
	botError "QuizBot/pkg/error"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

func (m *MongoRepository) AddQuiz(quiz *entity.Quiz) error {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Quiz.Name)
	_, err := coll.InsertOne(context.TODO(), quiz)
	if err != nil {
		return botError.NewBotError(err, m.Config.BotProblem)
	}
	return nil
}

func (m *MongoRepository) UpdateQuiz(chatID int64, isCorrectAns bool) error {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Quiz.Name)
	filter := bson.D{{m.Config.Quiz.ChatIDField, chatID}}
	incCurrent := bson.D{{"$inc", bson.D{{m.Config.Quiz.TaskIteratorField, 1}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, incCurrent)
	if err != nil {
		return botError.NewBotError(err, m.Config.BotProblem)
	}
	if isCorrectAns {
		incCorrect := bson.D{{"$inc", bson.D{{m.Config.Quiz.CorrectTasksField, 1}}}}
		_, err = coll.UpdateOne(context.TODO(), filter, incCorrect)

	}
	if err != nil {
		return botError.NewBotError(err, m.Config.BotProblem)
	}
	return nil
}

func (m *MongoRepository) GetCurrentTask(chatID int64) (*entity.Task, int, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Quiz.Name)
	filter := bson.D{{m.Config.Quiz.ChatIDField, chatID}}

	var quiz entity.Quiz

	err := coll.FindOne(context.TODO(), filter).Decode(&quiz)
	if err != nil {
		return nil, -1, botError.NewBotError(err, m.Config.BotProblem)

	}

	if quiz.CurrentTaskIter >= len(quiz.Tasks) {
		return nil, -1, botError.NewBotError(err, m.Config.BotProblem)
	}

	if quiz.CurrentTaskIter < 0 {
		return nil, -1, botError.NewBotError(err, m.Config.BotProblem)
	}

	return &quiz.Tasks[quiz.CurrentTaskIter], quiz.CurrentTaskIter, nil
}

func (m *MongoRepository) DeleteQuiz(chatID int64) (bool, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Quiz.Name)
	filter := bson.D{{m.Config.Quiz.ChatIDField, chatID}}

	c, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return false, botError.NewBotError(err, m.Config.BotProblem)
	}

	return c.DeletedCount > 0, nil
}

func (m *MongoRepository) GetQuiz(chatID int64) (*entity.Quiz, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Quiz.Name)
	filter := bson.D{{m.Config.Quiz.ChatIDField, chatID}}

	var quiz entity.Quiz
	err := coll.FindOne(context.TODO(), filter).Decode(&quiz)
	if err != nil {
		return nil, botError.NewBotError(err, m.Config.BotProblem)
	}

	return &quiz, err
}

func (m *MongoRepository) InitEmptyQuiz(chatID int64) error {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Quiz.Name)

	quiz := entity.Quiz{
		ChatID:          chatID,
		CurrentTaskIter: 0,
		TaskAmount:      0,
		CorrectTasks:    0,
		Sphere:          "",
		Section:         "",
		Tasks:           []entity.Task{},
		Tags:            []string{},
		Difficulty:      "",
	}
	_, err := coll.InsertOne(context.TODO(), &quiz)
	if err != nil {
		return botError.NewBotError(err, m.Config.BotProblem)
	}
	return nil
}

func (m *MongoRepository) UpdateQuizSettings(chatID int64, key, value string) error {
	if key == m.Config.Quiz.NoneValue {
		return nil
	}

	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Quiz.Name)

	filter := bson.D{{m.Config.Quiz.ChatIDField, chatID}}
	update := bson.M{"$set": bson.M{key: value}}
	if key == m.Config.Quiz.TaskAmountField {
		val, err := strconv.Atoi(value)
		if err != nil {
			return botError.NewBotError(err, m.Config.BotProblem)
		}
		update = bson.M{"$set": bson.M{m.Config.Quiz.TaskAmountField: val}}
	}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return botError.NewBotError(err, m.Config.BotProblem)
	}

	return nil
}

func (m *MongoRepository) SetQuizTasks(chatID int64, tasks []entity.Task) error {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Quiz.Name)

	filter := bson.D{{m.Config.Quiz.ChatIDField, chatID}}
	update := bson.M{"$set": bson.M{m.Config.Quiz.TasksField: tasks}}

	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return botError.NewBotError(err, m.Config.BotProblem)
	}
	return nil
}

func (m *MongoRepository) FillQuizWithTasks(chatID int64, tasks []entity.Task) error {
	quiz, err := m.GetQuiz(chatID)
	if err != nil {
		return err
	}

	err = m.SetQuizTasks(chatID, tasks)
	if err != nil {
		return err
	}

	if len(tasks) < quiz.TaskAmount {
		err = m.UpdateQuizSettings(chatID, m.Config.Quiz.TaskAmountField, strconv.Itoa(len(tasks)))
		if err != nil {
			return err
		}
	}

	if len(tasks) == 0 {
		return botError.NewBotError(errors.New("no files found according to provided quiz settings"), m.Config.NoFilesFound)
	}
	return nil
}
