package postgresDB

import (
	"QuizBot/pkg/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	DriverName = "postgres"
)

type PostgresRepository struct {
	DB     *sqlx.DB
	Config *config.PostgresConfig
}

func NewPostgresRepository(DB *sqlx.DB, cfg *config.PostgresConfig) *PostgresRepository {
	return &PostgresRepository{
		DB:     DB,
		Config: cfg,
	}
}

func InitPostgresRepository(cfg *config.PostgresConfig) (*PostgresRepository, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sqlx.Open(DriverName, conn)
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
	return NewPostgresRepository(db, cfg), nil
}
