package main

import "fmt"

func main() {
	var qty int

	fmt.Scan(&qty)
	nums := make([]int, qty)

	for i := 0; i < qty; i++ {
		var num int
		fmt.Scan(&num)
		nums[i] = num
	}

	var f, l, res int
	fmt.Scan(&f, &l)
	f -= 1
	l -= 1

	for i, v := range nums {
		if i >= f && i <= l {
			res += v
		}
	}

	fmt.Println(res)
}
