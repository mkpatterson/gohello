package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"John", "Dillinger", 27}
	p2 := person{"Terry", "Flaps", 31}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
