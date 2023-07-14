package mongodb

import (
	"QuizBot/internal/entity"
	botError "QuizBot/internal/error"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

func (m *MongoRepository) GetTasksForQuiz(quiz *entity.Quiz) ([]entity.Task, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Task.Name)

	filter := m.GenerateMongoFilter([]string{
		m.Config.Task.SphereField, m.Config.Task.SectionField, m.Config.Task.DifficultyField,
	},
		[]string{quiz.Sphere, quiz.Section, quiz.Difficulty})
	pipeline := []bson.M{
		{"$match": filter},
		{"$sample": bson.M{"size": quiz.TaskAmount}},
	}

	cursor, err := coll.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, botError.NewBotError(err, m.Config.NoFilesFound)
	}

	var tasks []entity.Task

	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, botError.NewBotError(err, m.Config.NoFilesFound)
	}

	return tasks, nil
}

func (m *MongoRepository) GetAllRequiredTasks(filter bson.M) ([]entity.Task, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Task.Name)
	cur, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, botError.NewBotError(err, m.Config.NoFilesFound)
	}
	defer func() {
		_ = cur.Close(context.Background())
	}()

	var tasks []entity.Task

	err = cur.All(context.Background(), &tasks)
	if err != nil {
		return nil, botError.NewBotError(err, m.Config.NoFilesFound)
	}
	return tasks, nil
}

func (m *MongoRepository) GenerateMongoFilter(keys, values []string) bson.M {
	filter := bson.M{}

	for i := 0; i < len(keys) && i < len(values); i++ {
		if values[i] == m.Config.Quiz.AllValue {
			continue
		}
		key := keys[i]
		value := values[i]
		filter[key] = value
	}

	return filter
}

func (m *MongoRepository) GetAllSpheres() ([]string, error) {
	empty := make([]string, 0)
	filter := m.GenerateMongoFilter(empty, empty)

	tasks, err := m.GetAllRequiredTasks(filter)
	if err != nil {
		return nil, err
	}

	unique := make(map[string]bool)
	for _, t := range tasks {
		for _, s := range t.Sphere {
			unique[s] = true
		}
	}

	spheres := make([]string, 0)
	for s := range unique {
		spheres = append(spheres, s)
	}
	return spheres, nil
}

func (m *MongoRepository) GetAllSectionsOfSphere(sphere string) ([]string, error) {
	keys := []string{m.Config.Task.SphereField}
	values := []string{sphere}
	filter := m.GenerateMongoFilter(keys, values)

	tasks, err := m.GetAllRequiredTasks(filter)
	if err != nil {
		return nil, err
	}

	unique := make(map[string]bool)
	for _, t := range tasks {
		for _, s := range t.Section {
			unique[s] = true
		}
	}

	sections := make([]string, 0)
	for s := range unique {
		sections = append(sections, s)
	}
	return sections, nil
}

func (m *MongoRepository) GetAllDifficultiesOfSelected(sphere string, section string) ([]string, error) {
	keys := []string{m.Config.Task.SphereField, m.Config.Task.SectionField}
	values := []string{sphere, section}
	filter := m.GenerateMongoFilter(keys, values)

	tasks, err := m.GetAllRequiredTasks(filter)
	if err != nil {
		return nil, err
	}

	unique := make(map[string]bool)
	for _, t := range tasks {
		unique[t.Difficulty] = true
	}

	difficulty := make([]string, 0)
	for s := range unique {
		difficulty = append(difficulty, s)
	}
	return difficulty, nil
}

func (m *MongoRepository) IsValidSphere(req string) (bool, error) {
	spheres, err := m.GetAllSpheres()
	if err != nil {
		return false, err
	}
	spheres = append(spheres, m.Config.Quiz.AllValue)
	return isObtain(spheres, req), nil
}

func (m *MongoRepository) IsValidSection(req, sphere string) (bool, error) {
	sections, err := m.GetAllSectionsOfSphere(sphere)
	if err != nil {
		return false, err
	}
	sections = append(sections, m.Config.Quiz.AllValue)
	return isObtain(sections, req), nil
}

func (m *MongoRepository) IsValidDifficulty(req, sphere, section string) (bool, error) {
	diff, err := m.GetAllDifficultiesOfSelected(sphere, section)
	if err != nil {
		return false, err
	}
	diff = append(diff, m.Config.Quiz.AllValue)
	return isObtain(diff, req), nil
}

func (m *MongoRepository) IsValid(req, typeParam string, quiz *entity.Quiz) (bool, error) {
	switch typeParam {
	case m.Config.Quiz.SphereField:
		return m.IsValidSphere(req)
	case m.Config.Quiz.SectionField:
		return m.IsValidSection(req, quiz.Sphere)
	case m.Config.Quiz.DifficultyField:
		return m.IsValidDifficulty(req, quiz.Sphere, quiz.Section)
	case m.Config.Quiz.TaskAmountField:
		return isValidAmount(req), nil
	case m.Config.Quiz.NoneValue:
		return true, nil
	}
	return false, botError.NewBotError(errors.New("invalid type to check validation"), m.Config.InvalidOption)
}

func isValidAmount(req string) bool {
	n, err := strconv.Atoi(req)
	if err != nil {
		return false
	}
	return n > 0
}
