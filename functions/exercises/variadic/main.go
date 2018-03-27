package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	greatest := max(33, 12, 66, 44, 7457, 8, 38, 2, 35, 36, 55, 7346, 7365, 8, 3578, 357, 8, 357, 235, 1345, 1023)
	fmt.Println(greatest)
}
