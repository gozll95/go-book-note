package main

import "fmt"

func main() {
	ch := make(chan int)
	go shower(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func shower(c chan int) {
	for {
		j := <-c
		fmt.Printf("%d\n", j)
	}
}

/*
其中一个麻烦是 goroutine 在 main.main() 结束的时候，没有进行清理。
更糟的是，由于 main.main() 和 main.shower() 的竞争关系，不是所有数字都被打印了。本应该打印到 9，但是有时只打印到 8。添加第二个退出 channel，可以解决这两个问题。试试吧。
*/
