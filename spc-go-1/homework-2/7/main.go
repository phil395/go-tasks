package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		input float64
		min   float64
		max   float64
		count int
	)

	for {
		fmt.Scan(&input)

		if input < 0 {
			fmt.Printf("%d\n%.1f %.1f\n", count, min, max)
			break
		}

		if (input < 100) || (input > 140) {
			continue
		}

		count++

		if min == 0 {
			min = input
		}

		if max == 0 {
			max = input
		}

		min = math.Min(min, input)
		max = math.Max(max, input)
	}
}
