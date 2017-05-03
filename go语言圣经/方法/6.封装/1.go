package main

import (
	"fmt"
)

func main() {
	a := [2]int{1, 2}
	var b []int
	b = a[:0]
	fmt.Println(b)
}
