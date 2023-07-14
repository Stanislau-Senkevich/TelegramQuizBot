package config

type PostgresConfig struct {
	DBHost     string `mapstructure:"postgres_host"`
	DBPort     string `mapstructure:"postgres_port"`
	DBName     string `mapstructure:"postgres_db_name"`
	DBUser     string `mapstructure:"postgres_user"`
	DBPassword string
	Stages
	Errors
}

type Stages struct {
	StartStage          string `mapstructure:"start"`
	SphereStage         string `mapstructure:"sphere"`
	SectionStage        string `mapstructure:"section"`
	DifficultyStage     string `mapstructure:"difficulty"`
	AmountStage         string `mapstructure:"amount"`
	QuizStage           string `mapstructure:"quiz"`
	StatisticsStage     string `mapstructure:"statistics"`
	FullStatisticsStage string `mapstructure:"full_statistics"`
}
