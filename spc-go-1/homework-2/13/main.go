package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		minX = math.MaxFloat64
		maxX = -math.MaxFloat64
		minY = math.MaxFloat64
		maxY = -math.MaxFloat64
	)

	for i := 0; i < 4; i++ {
		var x, y float64
		fmt.Scan(&x)
		fmt.Scan(&y)

		minX = math.Min(minX, x)
		maxX = math.Max(maxX, x)
		minY = math.Min(minY, y)
		maxY = math.Max(maxY, y)
	}

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var x, y float64
		fmt.Scan(&x)
		fmt.Scan(&y)

		if x >= minX && x <= maxX && y >= minY && y <= maxY {
			fmt.Printf("Точка с координатами %.0f,%.0f принадлежит исследуемому квадрату\n", x, y)
			continue
		}
		fmt.Printf("Точка с координатами %.0f,%.0f не принадлежит исследуемому квадрату\n", x, y)
	}

}
