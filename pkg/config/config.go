package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string
	Mongo         MongoConfig
	Postgres      PostgresConfig
	Messages
}

func InitConfig() (*Config, error) {
	viper.SetConfigFile("configs/config.yml")
	var cfg Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("mongo_config", &cfg.Mongo); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("postgres_config", &cfg.Postgres); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("responses", &cfg.Responses); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("errors", &cfg.Errors); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("errors", &cfg.Mongo.Errors); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("errors", &cfg.Postgres.Errors); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("buttons", &cfg.Buttons); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("stages", &cfg.Stages); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("stages", &cfg.Postgres.Stages); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("task_collection", &cfg.Mongo.Task); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("quiz_collection", &cfg.Mongo.Quiz); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("stats_collection", &cfg.Mongo.Stats); err != nil {
		return nil, err
	}

	cfg.AllButton = cfg.Mongo.Quiz.AllValue

	err := parseEnv(&cfg)

	return &cfg, err
}

func parseEnv(cfg *Config) error {
	//err := gotenv.Load()
	//if err != nil {
	//	return err
	//}

	if err := viper.BindEnv("telegram_token"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo_user"); err != nil {
		return err
	}

	if err := viper.BindEnv("mongo_password"); err != nil {
		return err
	}

	if err := viper.BindEnv("postgres_password"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("telegram_token")
	cfg.Mongo.User = viper.GetString("mongo_user")
	cfg.Mongo.Password = viper.GetString("mongo_password")
	cfg.Postgres.DBPassword = viper.GetString("postgres_password")

	return nil
}
