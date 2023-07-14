package entity

type Task struct {
	Question         string   `bson:"question"`
	Options          []string `bson:"options"`
	CorrectOptionIDs []int    `bson:"correct_answer_id"`
	Sphere           []string `bson:"sphere"`
	Section          []string `bson:"section"`
	Tags             []string `bson:"tags"`
	Difficulty       string   `bson:"difficulty"`
}
