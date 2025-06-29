package main

import (
	"fmt"
	"time"
)

type Student1 struct {
	Name      string
	BirthYear int
	Course    int
	AvgGrade  float64
}

func (s Student1) GetAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student1) GetStatus() string {
	switch {
	case s.AvgGrade >= 4.5:
		return "отличник"
	case s.AvgGrade >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func (s *Student1) Promote() {
	s.Course++
}

func main() {
	student := Student1{
		Name:      "Иван Петров",
		BirthYear: 2000,
		Course:    2,
		AvgGrade:  4.7,
	}

	fmt.Printf("%s:\n", student.Name)
	fmt.Printf("Возраст: %d лет\n", student.GetAge())
	fmt.Printf("Статус: %s\n", student.GetStatus())

	student.Promote()
	fmt.Printf("Новый курс: %d\n", student.Course)
}
