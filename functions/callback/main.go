package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{2, 4, 9, 7, 1, 3}, func(n int) {
		fmt.Println(n)
	})
}
