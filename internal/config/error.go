package config

type Errors struct {
	InvalidOption string `mapstructure:"invalid_option"`
	InvalidTime   string `mapstructure:"invalid_time"`
	UnknownError  string `mapstructure:"unknown_error"`
	NoFilesFound  string `mapstructure:"no_files"`
	StopQuiz      string `mapstructure:"stop_quiz"`
	BotProblem    string `mapstructure:"bot_problem"`
	SaveQuizError string `mapstructure:"save_quiz"`
}
