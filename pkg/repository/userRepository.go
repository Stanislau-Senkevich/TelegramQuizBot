package repository

type UserRepository interface {
	LogUser(chatID int64) error
	CheckUser(chatID int64) (bool, error)
	GetStage(chatID int64) (string, error)
	SetStage(chatID int64, stage string) error
}
