package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n    int
		dish string
	)

	fmt.Scan(&n)
	dishes := make(map[string]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&dish)
		dishes[dish] = 0
	}

	var qtyDays int
	fmt.Scan(&qtyDays)

	for i := 0; i < qtyDays; i++ {
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			fmt.Scan(&dish)
			dishes[dish]++
		}
	}

	result := []string{}

	for dish := range dishes {
		if dishes[dish] == 0 {
			result = append(result, dish)
		}
	}

	sort.Strings(result)

	for _, dish := range result {
		fmt.Println(dish)
	}
}
