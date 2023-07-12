package entity

import "fmt"

type Statistics struct {
	Spheres            map[string]bool
	Sections           map[string]bool
	Difficulties       map[string]bool
	TasksAmount        int
	CorrectTasksAmount int
	QuizAmount         int
}

func InitStatistics() *Statistics {
	return &Statistics{
		Spheres:            make(map[string]bool),
		Sections:           make(map[string]bool),
		Difficulties:       make(map[string]bool),
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
	for v := range s.Spheres {
		if v == "All" {
			continue
		}
		temp += " " + v + ","
	}
	ans += temp[:len(temp)-1]
	temp = ""
	ans += "\nSections:"
	for v := range s.Sections {
		if v == "All" {
			continue
		}
		temp += " " + v + ","
	}
	ans += temp[:len(temp)-1]
	temp = ""
	ans += "\nDifficulties:"
	for v := range s.Difficulties {
		if v == "All" {
			continue
		}
		temp += " " + v + ","
	}
	ans += temp[:len(temp)-1]
	return ans
}
