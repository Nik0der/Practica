package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func NewStudent(name string, age, course int, avgGrade float64) Student {
	return Student{
		Name:     name,
		Age:      age,
		Course:   course,
		AvgGrade: avgGrade,
	}
}

func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func (s *Student) UpdateGrade(newGrade float64) {
	s.AvgGrade = newGrade
	fmt.Printf("Новый средний балл %s: %.2f\n", s.Name, s.AvgGrade)
}

func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\nВозраст: %d\nКурс: %d\nСредний балл: %.2f\n",
		s.Name, s.Age, s.Course, s.AvgGrade)
}

func main() {
	student1 := NewStudent("Иван Иванов", 20, 2, 4.5)
	var student2 Student
	student2.Name = "Никита Каменский"
	student2.Age = 20
	student2.Course = 3
	student2.AvgGrade = 4.8

	student1.PrintInfo()
	student1.Promote()
	student1.UpdateGrade(4.7)
	fmt.Println()

	student2.PrintInfo()
}
