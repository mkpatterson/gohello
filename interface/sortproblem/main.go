package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	studyGroup := people{"Zeno", "Jon", "Al", "Jenny"}
	sort.Strings(studyGroup)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)

	s := []string{"Zeno", "Jon", "Al", "Jenny"}
	sort.Strings(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(" ")
	fmt.Println(s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(" ")
	fmt.Println(n)
}
