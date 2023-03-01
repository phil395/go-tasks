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
	result := ""

	for i, v := range chars {
		if i%2 != 0 {
			continue
		}
		result += fmt.Sprintf("%s%[1]s%[1]s", string(v))
	}
	fmt.Println(result)

}
