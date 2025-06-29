package main

import "fmt"

func removeElement(slice []string, index int) []string {

	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func main() {
	fruits := []string{"Яблоко", "Банан", "Апельсин", "Киви", "Манго"}
	fmt.Println("Исходный срез:", fruits)

	fruits = removeElement(fruits, 2)
	fmt.Println("После удаления:", fruits)

	fruits = removeElement(fruits, 0)
	fmt.Println("После удаления первого элемента:", fruits)

	fruits = removeElement(fruits, 10)
	fmt.Println("Попытка удалить несуществующий элемент:", fruits)
}
