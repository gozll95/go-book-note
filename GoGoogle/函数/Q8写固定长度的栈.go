/*
创建一个固定大小保存整数的栈。它无须超出限制的增长。定义 push 函数——将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（LIFO）的。
*/

/*
package main

import (
	"fmt"
)

type s1 []int

func (s s1) pushZhan(x int) {
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			s[i] = x
			break
		}
	}

	for index, v := range s {
		fmt.Printf("%v:%p", index, v)
	}
}

func (s s1) popZhan() {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 0 {
			s.Remove()
			break
		}
	}
	for index, v := range s {
		fmt.Printf("%v:%p", index, v)
	}
}

func main() {
	var s []int
	a := [100]int{1: 99}
	s = a[1:10]
	s.pushZhan(1)
	s.pushZhan(1)
}

*/

package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	// if s.i+1 > 9 {
	// 	return
	// }

	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	fmt.Println(s.data[s.i])
	return s.data[s.i]
}

func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main() {
	s := new(stack)
	s.push(25)
	fmt.Printf("stack %v\n", s)
	s.push(10)
	fmt.Printf("stack %v\n", s)
	s.pop()
	fmt.Printf("stack %v\n", s)
}
