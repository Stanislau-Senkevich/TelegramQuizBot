package botError

type BotError struct {
	Err    error
	BotMsg string
}

func NewBotError(err error, msg string) *BotError {
	return &BotError{
		Err:    err,
		BotMsg: msg,
	}
}

func (e *BotError) Error() string {
	return e.Err.Error() + "\n" + e.BotMsg
}
