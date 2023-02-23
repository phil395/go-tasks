package main

import (
	"fmt"
)

func main() {
	var a, b float32
	fmt.Scan(&a, &b)

	if int(a+b)%2 == 0 {
		fmt.Println("ЧЁТНОЕ")
		return
	}

	fmt.Println("НЕЧЁТНОЕ")
}
