package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("launch!")
}

/*
time.Ticker
go的time和ticket的调用

或者叫timmer internal和其他语言的开发思路完全不一样。

其他语言，多是注册回调函数，定时，时间到了调用回调。

go是 通过 chan

的阻塞实现的。

调用的地方，读取chan

定时，时间到，向chan写入值，阻塞解除，调用函数
*/
