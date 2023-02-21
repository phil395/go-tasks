package main

import "fmt"

func main() {
	var (
		firstName, lastName string
		age                 int
	)
	fmt.Scan(&firstName, &lastName, &age)
	fmt.Printf("Имя: %s, Фамилия: %s , Возраст: %d . Студент BPS\n", firstName, lastName, age)
}
