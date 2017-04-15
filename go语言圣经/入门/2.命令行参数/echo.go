package main

import (
	"fmt"
	"os"
	"strings"
)

/*
func main() {
	var s string
	for _, v := range os.Args[1:] {
		s = s + " " + v
	}
	fmt.Println(s)
}
*/
//改进
/*
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
*/

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
