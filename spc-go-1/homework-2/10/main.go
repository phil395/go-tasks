package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Scan(&num)

	for i := 0 - num; i <= num; i++ {
		fmt.Printf("Квадрат числа %d равен %d\n", i, i*i)
	}

}
