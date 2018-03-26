package main

import "fmt"

func main() {
	for i := 1; i < 500; i++ {
		fmt.Print(i, " - ", string(i), " - ", []byte(string(i)), " - ")
		fmt.Printf("%b\n", i)
	}
}
