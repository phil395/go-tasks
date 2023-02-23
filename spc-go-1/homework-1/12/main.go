package main

import "fmt"

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)

	if (a == "раз" || a == "один") && b == "два" && c == "три" {
		fmt.Println("ОК")
		return
	}

	if a == "1" && b == "2" && c == "3" {
		fmt.Println("ОК")
		return
	}

	fmt.Println("НЕ ПРАВИЛЬНО")
}
