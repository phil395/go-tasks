package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lastChar := ' '

	for scanner.Scan() {
		word := scanner.Text()
		runes := []rune(word)

		if lastChar == ' ' || runes[0] == lastChar {
			char := runes[len(runes)-1]
			if char == 'ь' {
				char = runes[len(runes)-2]
			}
			lastChar = char
			continue
		}

		fmt.Println(word)
		break

	}
}
