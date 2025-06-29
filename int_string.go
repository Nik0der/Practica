package main

import "fmt"

func main() {

	var age int = 20
	var name string = "Никита"
	var isStudent bool = true
	var height float64 = 177

	fmt.Printf("Возраст: %d\n", age)
	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Студент: %t\n", isStudent)
	fmt.Printf("Рост: %.2f см\n", height)
}
