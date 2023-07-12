package config

type MongoConfig struct {
	User             string
	Password         string
	DBName           string `mapstructure:"db_name"`
	ConnectionString string `mapstructure:"conn_string"`
	Task             TaskCollection
	Quiz             QuizCollection
	Stats            StatisticsCollection
	Errors
}

type TaskCollection struct {
	Name                 string `mapstructure:"name"`
	QuestionField        string `mapstructure:"question_field"`
	OptionsField         string `mapstructure:"options_field"`
	CorrectOptionIDField string `mapstructure:"correct_option_id_field"`
	SphereField          string `mapstructure:"sphere_field"`
	SectionField         string `mapstructure:"section_field"`
	DifficultyField      string `mapstructure:"difficulty_field"`
	TagsField            string `mapstructure:"tags_field"`
}

type QuizCollection struct {
	Name              string `mapstructure:"name"`
	ChatIDField       string `mapstructure:"chat_id_field"`
	SphereField       string `mapstructure:"sphere_field"`
	SectionField      string `mapstructure:"section_field"`
	DifficultyField   string `mapstructure:"difficulty_field"`
	TagsField         string `mapstructure:"tags_field"`
	TaskAmountField   string `mapstructure:"task_amount_field"`
	TaskIteratorField string `mapstructure:"task_iter_field"`
	CorrectTasksField string `mapstructure:"correct_tasks_field"`
	TasksField        string `mapstructure:"tasks_field"`
	NoneValue         string `mapstructure:"none_value"`
	AllValue          string `mapstructure:"all_value"`
}

type StatisticsCollection struct {
	Name              string `mapstructure:"name"`
	ChatIDField       string `mapstructure:"chat_id_field"`
	SphereField       string `mapstructure:"sphere_field"`
	SectionField      string `mapstructure:"section_field"`
	DifficultyField   string `mapstructure:"difficulty_field"`
	TaskAmountField   string `mapstructure:"task_amount_field"`
	CorrectTasksField string `mapstructure:"correct_tasks_field"`
	DateField         string `mapstructure:"date_field"`
}
