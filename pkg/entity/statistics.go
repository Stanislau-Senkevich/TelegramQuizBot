package entity

import "fmt"

type Statistics struct {
	Spheres            map[string]int
	Sections           map[string]int
	Difficulties       map[string]int
	TasksAmount        int
	CorrectTasksAmount int
	QuizAmount         int
}

func InitStatistics() *Statistics {
	return &Statistics{
		Spheres:            make(map[string]int),
		Sections:           make(map[string]int),
		Difficulties:       make(map[string]int),
		TasksAmount:        0,
		CorrectTasksAmount: 0,
		QuizAmount:         0,
	}
}

func (s *Statistics) Print() string {
	ans := "Your statistics\n"
	ans += fmt.Sprintf("Quizzes passed: %d\nSummary results: %d/%d\n", s.QuizAmount, s.CorrectTasksAmount, s.TasksAmount)
	ans += "Spheres:"
	temp := ""
	for v, am := range s.Spheres {
		if v == "All" {
			temp += " Mixed"
		} else {
			temp += " " + v
		}
		if am > 1 {
			temp += fmt.Sprintf("(x%d)", am)
		}
		temp += ","
	}
	ans += temp[:len(temp)-1]
	temp = ""
	ans += "\nSections:"
	for v, am := range s.Sections {
		if v == "All" {
			temp += " Mixed"
		} else {
			temp += " " + v
		}
		if am > 1 {
			temp += fmt.Sprintf("(x%d)", am)
		}
		temp += ","
	}
	ans += temp[:len(temp)-1]
	temp = ""
	ans += "\nDifficulties:"
	for v, am := range s.Difficulties {
		if v == "All" {
			temp += " Mixed"
		} else {
			temp += " " + v
		}
		if am > 1 {
			temp += fmt.Sprintf("(x%d)", am)
		}
		temp += ","
	}
	ans += temp[:len(temp)-1]
	return ans
}
