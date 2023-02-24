package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	d := b*b - 4*a*c

	if (a == 0) || (d == 0) {
		fmt.Println("один корень")
		return
	}

	if d < 0 {
		fmt.Println("корней нет")
		return
	}

	fmt.Println("два корня")
}
