package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	input := scanner.Text()
	chars := []rune(input)
	fChar := chars[0]
	lChar := chars[len(chars)-1]

	if (fChar == 'Д' || fChar == 'д') && (lChar == 'а' || lChar == 'А') {
		fmt.Println("СОГЛАСЕН")
		return
	}
	fmt.Println("НЕ СОГЛАСЕН")
}
