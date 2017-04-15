package main

import (
	"fmt"
)

// func shower(c chan int, quit chan bool, i int) {
// 	fmt.Println("i is", i)
// 	for {
// 		select {
// 		case c <- i:
// 			fmt.Println(i)
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}

// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan bool)

// 	for i := 1; i < 10; i++ {
// 		go shower(c, quit, i)
// 	}

// 	for i := 1; i < 10; i++ {
// 		<-c
// 	}
// 	quit <- true

// }

func shower(c chan int, quit chan bool) {
	for {
		select {
		case j := <-c:
			fmt.Printf("%d\n", j)
		case <-quit:
			break
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go shower(ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- true
}
