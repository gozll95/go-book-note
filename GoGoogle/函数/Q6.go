/*
整数顺序

编写函数，返回其（两个）参数正确的（自然）数字顺序：
*/

package main

import (
	"fmt"
)

func sort(x, y int) (a, b int) {
	if x > y {
		t := x
		x = y
		y = t
	}
	return x, y
}

func main() {
	a := 2
	b := 1
	a, b = sort(a, b)
	fmt.Println("after sort", a, b)
}

/*

func order(a,b int)(int,int){
	if a>b{
		return b,a
	}
	return a,b
}
*/
