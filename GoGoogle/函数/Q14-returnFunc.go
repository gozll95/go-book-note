package main

import (
	"fmt"
)

// func plusTwo(y int) (f func(x int) int) {
// 	return func(x int) int {
// 		return x + y
// 	}
// }

func plusX(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	p := plusTwo(2)
	fmt.Printf("%v\n", p(3))
}
