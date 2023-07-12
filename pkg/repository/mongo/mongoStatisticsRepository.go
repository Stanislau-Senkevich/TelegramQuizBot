package mongoDB

import (
	"QuizBot/pkg/entity"
	botError "QuizBot/pkg/error"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
	"time"
)

func (m *MongoRepository) SaveResults(quiz *entity.Quiz) error {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Stats.Name)

	statQuiz := quiz.GetStatVersion()
	_, err := coll.InsertOne(context.TODO(), statQuiz)
	if err != nil {
		return botError.NewBotError(err, m.Config.SaveQuizError)
	}
	return nil
}

func (m *MongoRepository) GetStatistics(chatID int64, timeValue time.Duration) ([]entity.StatQuiz, error) {
	coll := m.DB.Database(m.Config.DBName).Collection(m.Config.Stats.Name)

	filter := bson.M{
		m.Config.Stats.DateField: bson.M{
			"$gte": time.Now().Add(-timeValue),
		},
		m.Config.Stats.ChatIDField: chatID,
	}

	var stats []entity.StatQuiz

	cur, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, botError.NewBotError(err, m.Config.BotProblem)
	}

	err = cur.All(context.Background(), &stats)
	if err != nil {
		return nil, botError.NewBotError(err, m.Config.BotProblem)
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Date.After(stats[j].Date)
	})
	return stats, nil
}
