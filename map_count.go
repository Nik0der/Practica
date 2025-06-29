package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go это язык программирования Go разработан в Google язык Go простой и эффективный"

	words := strings.Fields(text)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[strings.ToLower(word)]++
	}

	fmt.Println("Частота слов:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}

	mostFrequent := ""
	maxCount := 0
	for word, count := range wordCount {
		if count > maxCount {
			mostFrequent = word
			maxCount = count
		}
	}
	fmt.Printf("\nСамое частое слово: '%s' (%d раз)\n", mostFrequent, maxCount)
}
