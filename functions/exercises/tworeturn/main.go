package main

import "fmt"

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	for i := 0; i <= 100; i++ {
		i, even := half(i)
		fmt.Println(i, even)
	}
}
