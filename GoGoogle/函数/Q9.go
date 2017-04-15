//缓冲式channel
//close 显示关闭channel
//可以通过语法v,ok :=<-ch,测试 channel是否被关闭

/*
package main

import (
	"fmt"
)

func fibonnacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonnacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

*/

package main

import (
	"fmt"
)

func fibonnacci(value int) []int {
	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x
}

func main() {
	for _, term := range fibonnacci(10) {
		fmt.Printf("%v ", term)
	}
}
