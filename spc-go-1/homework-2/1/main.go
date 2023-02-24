package main

import (
	"fmt"
	"unicode/utf8"
)

const coastPerChar = 23

func main() {
	var input string
	fmt.Scan(&input)

	qtyChar := utf8.RuneCountInString(input)

	cost := qtyChar * coastPerChar

	fmt.Printf("%d р. %d коп.", cost/100, cost-(cost/100)*100)
}
