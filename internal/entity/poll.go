package entity

type Poll struct {
	PollID    string `db:"poll_id"`
	ChatID    int64  `db:"chat_id"`
	MessageID int    `db:"message_id"`
}
