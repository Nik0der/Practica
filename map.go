package main

import "fmt"

func main() {

	grades := make(map[string]int)

	addGrade := func(name string, grade int) {
		grades[name] = grade
		fmt.Printf("Добавлена оценка для %s: %d\n", name, grade)
	}

	getGrade := func(name string) (int, bool) {
		grade, exists := grades[name]
		return grade, exists
	}

	removeStudent := func(name string) {
		delete(grades, name)
		fmt.Printf("Студент %s удален\n", name)
	}

	addGrade("Никита", 5)
	addGrade("Иван", 4)
	addGrade("Александр", 3)

	if grade, exists := getGrade("Мария"); exists {
		fmt.Printf("Оценка Ивана: %d\n", grade)
	} else {
		fmt.Println("Студент не найден")
	}

	removeStudent("Александр")

	fmt.Println("\nТекущие оценки:")
	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}
}
