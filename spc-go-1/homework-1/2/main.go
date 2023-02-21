package main

import "fmt"

func main() {
	word1 := "имя"
	word2 := "твое"
	word3 := "мне"
	word4 := "знакомо"

	fmt.Println(word4, word3, word2, word1) // знакомо мне твое имя
	fmt.Println(word3, word1, word4, word2) // мне имя знакомо твое
	fmt.Println(word2, word4, word1, word3) // твое знакомо имя мне
}
