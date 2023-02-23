package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var login, email string
	fmt.Scan(&login, &email)

	loginValid := utf8.RuneCountInString(login) >= 10 && !strings.Contains(login, "@")

	if !loginValid {
		fmt.Println("Некорректный логин")
		return
	}

	emailValid := strings.Contains(email, "@") && strings.Contains(email, ".")

	if !emailValid {
		fmt.Println("Некорректная почта")
		return
	}

	fmt.Println("ОК")
}
