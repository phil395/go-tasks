package main

import "fmt"

func main() {
	var drink, meal, subMeal, time string

	fmt.Scan(&drink, &meal, &subMeal, &time)
	fmt.Printf("I wanna some %q, my favorite meal - %q. Give me also %q. I will come soon... %q\n", drink, meal, subMeal, time)
}
