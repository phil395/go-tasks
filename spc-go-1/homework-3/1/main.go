package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		a, b  int
		names = map[string]int{}
		res   []string
	)
	fmt.Scan(&a, &b)

	scn := bufio.NewScanner(os.Stdin)

	for i := 0; i < a; i++ {
		if scn.Scan() != true {
			fmt.Println("error")
		}
		name := scn.Text()
		names[name] = 1
	}

	for i := 0; i < b; i++ {
		if scn.Scan() != true {
			fmt.Println("error")
		}
		name := scn.Text()
		r := "НЕТ В НАЛИЧИИ"
		if names[name] == 1 {
			r = "ЕСТЬ"
		}
		res = append(res, r)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}
