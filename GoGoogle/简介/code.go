package main

import "fmt"

var P = fmt.Println

func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	P(i)
}

func main() {
	rec(3)
}
