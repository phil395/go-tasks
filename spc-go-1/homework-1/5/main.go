package main

import "fmt"

func main() {

	words := make([]string, 5, 5)

	for i := range words {
		var word string
		fmt.Scan(&word)
		words[i] = word
	}

	for i := len(words) - 2; i >= 0; i-- {
		fmt.Printf("%s - %s\n", words[i], words[4])
	}

}
