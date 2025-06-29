package main

import "fmt"

func main() {

	colors := make([]string, 0, 3)
	fmt.Printf("Исходный срез: %v, длина: %d, емкость: %d\n", colors, len(colors), cap(colors))

	colors = append(colors, "Красный")
	colors = append(colors, "Зеленый")
	colors = append(colors, "Синий")
	fmt.Printf("После добавления: %v, длина: %d, емкость: %d\n", colors, len(colors), cap(colors))

	colors = append(colors, "Желтый")
	colors = append(colors, "Фиолетовый")
	fmt.Printf("После расширения: %v, длина: %d, емкость: %d\n", colors, len(colors), cap(colors))

	fmt.Println("\nВсе элементы среза:")
	for i, color := range colors {
		fmt.Printf("%d: %s\n", i, color)
	}
}
