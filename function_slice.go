package main

import (
	"fmt"
	"sort"
)

func FindElement(slice []int, target int) (int, bool) {
	for index, value := range slice {
		if value == target {
			return index, true
		}
	}
	return -1, false
}

func SortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func FilterSlice(slice []int, condition func(int) bool) []int {
	var result []int
	for _, value := range slice {
		if condition(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}

	if index, found := FindElement(numbers, 8); found {
		fmt.Printf("Элемент 8 найден на позиции %d\n", index)
	} else {
		fmt.Println("Элемент не найден")
	}

	sorted := SortSlice(numbers)
	fmt.Println("Отсортированный срез:", sorted)

	evenNumbers := FilterSlice(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Четные числа:", evenNumbers)

	largeNumbers := FilterSlice(numbers, func(x int) bool {
		return x > 5
	})
	fmt.Println("Числа больше 5:", largeNumbers)
}
