package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "GoLang — современный язык программирования"

	length := len(text)
	fmt.Printf("Длина строки: %d символов\n", length)

	substring := "язык"
	contains := strings.Contains(text, substring)
	fmt.Printf("Содержит подстроку '%s': %t\n", substring, contains)

	upper := strings.ToUpper(text)
	lower := strings.ToLower(text)
	fmt.Println("Верхний регистр:", upper)
	fmt.Println("Нижний регистр:", lower)

	fmt.Println("Начинается с 'Go':", strings.HasPrefix(text, "Go"))
	fmt.Println("Заканчивается на 'я':", strings.HasSuffix(text, "я"))
	fmt.Println("Количество букв 'о':", strings.Count(text, "о"))
}
