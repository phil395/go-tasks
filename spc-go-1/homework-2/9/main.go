package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		num       int
		primeNums = []string{"1"}
	)

	fmt.Scan(&num)

	for i := 2; i <= num; i++ {
		if num%i == 0 {
			primeNums = append(primeNums, fmt.Sprint(i))
		}
	}

	fmt.Printf("%s\n", strings.Join(primeNums, " "))

	if len(primeNums) == 2 {
		fmt.Println("ACHTUNG")
	}
}
