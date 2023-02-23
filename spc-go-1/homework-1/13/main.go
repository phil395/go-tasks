package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)

	if (math.Abs(x1-x2) == 1) && (math.Abs(y1-y2) == 2) {
		fmt.Println("ДА")
		return
	}

	if (math.Abs(x1-x2) == 2) && (math.Abs(y1-y2) == 1) {
		fmt.Println("ДА")
		return
	}

	fmt.Println("НЕТ")
}
