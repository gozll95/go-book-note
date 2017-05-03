package main

import (
	"fmt"
)

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	f(3)
}

//当f(0)被调用时，发生panic异常，之前被延迟执行的的3个fmt.Printf被调用。程序中断执行后，panic信息和堆栈信息会被输出（下面是简化的输出）：
