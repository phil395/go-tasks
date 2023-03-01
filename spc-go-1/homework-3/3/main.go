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
			lastChar = runes[len(runes)-1]
			continue
		}

		fmt.Println(word)
		break

	}
}
