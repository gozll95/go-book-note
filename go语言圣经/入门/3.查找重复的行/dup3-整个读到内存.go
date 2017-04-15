//我们也可以一次性的把整个输入内容全部读到内存中，然后再把其分割为多行，然后再去处理这些行内的数据。下面的dup3这个例子就是以这种形式来进行操作的。这个例子引入了一个新函数ReadFile（从io/ioutil包提供），这个函数会把一个指定名字的文件内容一次性调入，之后我们用strings.Split函数把文件分割为多个子字符串，并存储到slice结构中。（Split函数是strings.Join的逆函数，Join函数之前提到过）

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//ReadFile函数返回byte类型的slice，这个slice必须被转换为string，之后才能够用strings.Split方法来进行处理
