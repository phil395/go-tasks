package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var qty int
	fmt.Scan(&qty)
	tasks := make([]string, qty)

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < qty; i++ {
		scanner.Scan()
		tasks[i] = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error in scanner")
	}

	fmt.Scan(&qty)
	selectedTasks := make([]string, qty)
	for i := 0; i < qty; i++ {
		var n int
		fmt.Scan(&n)
		if n > len(tasks) || n < 1 {
			fmt.Println("Incorrect task number")
		}
		selectedTasks[i] = tasks[n-1]
	}

	for _, task := range selectedTasks {
		fmt.Println(task)
	}
}
