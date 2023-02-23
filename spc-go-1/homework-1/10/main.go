package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	if strings.Contains(s, "чат") {
		fmt.Println("БОТ")
		return
	}
	fmt.Println("НЕ БОТ")
}
