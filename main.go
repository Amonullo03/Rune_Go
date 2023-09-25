package main

import (
	"fmt"
	"strings"
)

//режет длину текста до нужной длины (поддерживает русский и тадж языки)

//func main() {
//	str := truncateString("Amonullo", 2)
//	fmt.Println(str)
//}
//
//func truncateString(str string, length int) string {
//	runes := []rune(str)
//	if len(runes) <= length {
//		return str
//	}
//	str = ""
//	for i := 0; i <= length; i++ {
//		str = str + string(runes[i])
//	}
//	return str
//}

func main() {
	text := "Противоположная точка зрения подразумевает, что независимые государства, превозмогая сложившуюся непростую экономическую ситуацию, превращены в посмешище, хотя само их существование приносит несомненную пользу обществу."

	// Разбиваем текст на слова
	words := strings.Fields(text)
	line := ""
	for _, word := range words {
		if len(line) == 0 {
			line = word
		} else {
			line += " " + word
		}
		if len(line) >= 20 {
			fmt.Println(line)
			line = ""
		}
	}
	if len(line) > 0 {
		fmt.Println(line)
	}

	runes := []rune(text)
	for index, letter := range runes {
		fmt.Printf("index: %v letter: %c\n", index, letter)
	}
}
