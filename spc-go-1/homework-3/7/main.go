package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var (
		count    float64
		maxCount float64
	)

	for _, char := range input {
		if char != 'о' {
			count = 0
			continue
		}
		count++
		maxCount = math.Max(count, maxCount)
	}
	fmt.Println(maxCount)
}
