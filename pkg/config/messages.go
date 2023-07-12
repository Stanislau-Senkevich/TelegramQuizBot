package config

type Messages struct {
	Errors
	Responses
	Buttons
	Stages
}

type Responses struct {
	SelectSphere     string `mapstructure:"select_sphere"`
	SelectSection    string `mapstructure:"select_section"`
	SelectDifficulty string `mapstructure:"select_diff"`
	SelectAmount     string `mapstructure:"select_amount"`
	QuizStopped      string `mapstructure:"quiz_stop"`
	NoQuizToStop     string `mapstructure:"no_quiz_stop"`
	QuizStarted      string `mapstructure:"quiz_start"`
	QuizResults      string `mapstructure:"quiz_results"`
	QuizCorrect      string `mapstructure:"quiz_correct"`
	QuizWrong        string `mapstructure:"quiz_wrong"`
	StartReply       string `mapstructure:"start_reply"`
	AboutReply       string `mapstructure:"about_reply"`
	ContactsReply    string `mapstructure:"contacts_reply"`
	DurationReply    string `mapstructure:"duration_reply"`
	NoStatistics     string `mapstructure:"no_stats"`
	StatsPattern     string `mapstructure:"stats_pattern"`
}

type Buttons struct {
	AllButton        string
	StartCommand     string `mapstructure:"start"`
	StopButton       string `mapstructure:"stop"`
	QuizButton       string `mapstructure:"quiz"`
	StatisticsButton string `mapstructure:"statistics"`
	ContactsButton   string `mapstructure:"contacts"`
	AboutButton      string `mapstructure:"about"`
	Days7            string `mapstructure:"7_days"`
	Days14           string `mapstructure:"14_days"`
	Days30           string `mapstructure:"30_days"`
	Days90           string `mapstructure:"90_days"`
	AllTime          string `mapstructure:"all_time"`
}
