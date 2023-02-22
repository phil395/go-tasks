package main

import "fmt"

func main() {
	var width, height int
	fmt.Scan(&width, &height)
	fmt.Printf("Периметр: %d\n", 2*(width+height))
	fmt.Printf("Площадь: %d\n", width*height)
}
