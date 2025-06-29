package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Разбиваем строку на слова по пробелам"
	words := strings.Split(sentence, " ")

	fmt.Println("Исходное предложение:", sentence)
	fmt.Println("Слова:")
	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}

	complexSentence := "Раз, два, три!"
	fields := strings.FieldsFunc(complexSentence, func(r rune) bool {
		return r == ' ' || r == ',' || r == '!'
	})
	fmt.Println("\nСложное предложение:", complexSentence)
	fmt.Println("Слова (с обработкой знаков препинания):")
	for _, word := range fields {
		if word != "" {
			fmt.Println(word)
		}
	}
}
