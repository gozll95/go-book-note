package main

import (
	"fmt"
)

func Map(f func(int) int, a []int) []int {
	s := make([]int, len(a))
	for i, v := range a {
		s[i] = f(v)
	}
	return s
}

func f1(x int) int {
	return x * x
}

func main() {
	var s []int
	a := []int{1, 2, 3}
	s = Map(f1, a)
	// for _, v := range s {
	// 	fmt.Printf("%v ", v)
	// }
	fmt.Printf("%v ", s)
}
