package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	chars := []rune(input)
	// fmt.Println(chars)

	for i := range chars {
		if len(chars) <= 2 {
			break
		}

		if i%2 == 0 {
			chars = chars[2:]
		} else {
			chars = chars[:len(chars)-2]
		}

		fmt.Println(string(chars))
	}
}
