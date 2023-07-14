package entity

type User struct {
	ChatID int64  `db:"chat_id"`
	Stage  string `db:"stage"`
}
