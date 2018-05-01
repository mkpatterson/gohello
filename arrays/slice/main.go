package main

import "fmt"

func main() {
	sl := make([]int, 1, 5)

	fmt.Println("*-----------------------------------*")
	fmt.Println(sl)
	fmt.Println(len(sl))
	fmt.Println(cap(sl))
	fmt.Println("*-----------------------------------*")

	for i := 1; i < 50; i++ {
		sl = append(sl, i)
		fmt.Println("Len:", len(sl), "Cap:", cap(sl), "Val:", sl[i])
	}
}
