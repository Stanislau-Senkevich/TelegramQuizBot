package mongoDB

import (
	"QuizBot/pkg/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	DB     *mongo.Client
	Config *config.MongoConfig
}

func NewMongoRepository(DB *mongo.Client, cfg *config.MongoConfig) *MongoRepository {
	return &MongoRepository{
		DB:     DB,
		Config: cfg,
	}
}

func InitMongoRepository(cfg *config.MongoConfig) (*MongoRepository, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	conn := fmt.Sprintf(cfg.ConnectionString, cfg.User, cfg.Password)
	opts := options.Client().ApplyURI(conn).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	if err = client.Database(cfg.DBName).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, err
	}
	return NewMongoRepository(client, cfg), nil
}

func isObtain[T comparable](slice []T, subject T) bool {
	for _, v := range slice {
		if v == subject {
			return true
		}
	}
	return false
}
