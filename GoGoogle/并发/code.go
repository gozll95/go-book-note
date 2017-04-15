/*
package main

import (
	"fmt"
	"time"
)

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}

func main() {
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	// time.Sleep(5 * time.Second)
}
*/

/*
package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}

func main() {
	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	<-c
	<-c
	// time.Sleep(5 * time.Second)
}
*/

package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}

func main() {
	c = make(chan int, 2)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	i := 0
L:
	for {
		select {
		case <-c:
			i++
			fmt.Println(i)
			if i > 1 {

				break L
			}
		}
	}
	// time.Sleep(5 * time.Second)
}
