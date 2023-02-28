package main

import "fmt"

func main() {
	var n, res int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		if i%2 == 0 {
			res += v
		} else {
			res -= v
		}
	}

	fmt.Println(res)
}
