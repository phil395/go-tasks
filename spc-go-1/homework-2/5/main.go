package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var p1, p2 string

	for {
		fmt.Scan(&p1)
		fmt.Scan(&p2)

		if utf8.RuneCountInString(p1) < 8 {
			fmt.Println("Слишком короткий пароль!")
			continue
		}

		if strings.Contains(p1, "123") || strings.Contains(p1, "qwe") {
			fmt.Println("Слишком простой пароль!")
			continue
		}

		if p1 != p2 {
			fmt.Println("Введенные пароли различаются!")
			continue
		}

		fmt.Println("Ну наконец-то!")
		return
	}
}
