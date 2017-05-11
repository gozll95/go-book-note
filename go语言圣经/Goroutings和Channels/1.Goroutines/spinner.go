package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			//fmt.Printf("\r%c", r)
			fmt.Printf("\r%c", r)
			time.Sleep(delay)

		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

//\r：回车符，返回到这一行的开头，return的意思
//然后主函数返回。主函数返回时，所有的goroutine都会被直接打断，程序退出。
