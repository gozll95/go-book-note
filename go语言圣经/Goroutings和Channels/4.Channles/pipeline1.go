package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	//Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	//Printer(in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

/*
为了有限循环
改进1:

// Squarer
go func() {
    for {
        x, ok := <-naturals
        if !ok {
            break // channel was closed and drained
        }
        squares <- x * x
    }
    close(squares)
}()

使用range循环是上面处理模式的简洁语法，它依次从channel接收数据，当channel被关闭并且没有值可接收时跳出循环。

*/
