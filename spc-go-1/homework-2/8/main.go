package main

import "fmt"

func getNod(a int, b int) int {
	for b != 0 {
		a = a % b
		b, a = a, b
	}
	return a
}

func main() {
	var (
		qty, denominator, numerator int
	)

	fmt.Scan(&qty)
	pairs := make([][2]int, qty)

	for i := 0; i < qty; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)
		pairs[i][0] = a
		pairs[i][1] = b
		if i == 0 {
			denominator = b
			continue
		}
		denominator *= b
	}

	for _, pair := range pairs {
		numerator += pair[0] * denominator / pair[1]
	}

	nod := getNod(numerator, denominator)

	numerator /= nod
	denominator /= nod

	fmt.Printf("%d/%d\n", numerator, denominator)
}
