package main

import "fmt"

func LongestString(strings ...string) (string, int) {
	if len(strings) == 0 {
		return "", 0
	}

	longest := strings[0]
	for _, str := range strings[1:] {
		if len(str) > len(longest) {
			longest = str
		}
	}
	return longest, len(longest)
}

func main() {

	longest, length := LongestString(
		"Go",
		"Python",
		"JavaScript",
		"TypeScript",
		"Java",
		"C",
	)

	fmt.Printf("Самая длинная строка: '%s' (%d символов)\n", longest, length)

	languages := []string{"Ruby", "Rust", "Kotlin", "Swift"}
	longestFromSlice, _ := LongestString(languages...)
	fmt.Println("Самый длинный язык из среза:", longestFromSlice)
}
