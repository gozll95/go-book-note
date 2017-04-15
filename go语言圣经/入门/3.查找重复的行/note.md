文件拷贝、文件打印、文件搜索、文件排序、文件统计类的程序：
一个处理输入的循环，在每一个输入元素上执行计算处理，在处理的同事或者处理完成之后进行结果输出。


2）
#bufio package中最有用的是Scanner类型，可以简单实现接收输入，或者把输入打散成行或者单词；这个工具通常是处理行形式的输入最简单的方法了。

#scanner对象从程序的标准输入中读取内容。对input.Scanner的每一次调用都会调入一个新行，并且会自动将其行末的换行符去掉；结果用input.Text()得到。Scan方法在读到了新行的时候会返回true,而在没有新行被读入时，会返回false。

3)fmt.Printf 
%d          int变量
%x, %o, %b  分别为16进制，8进制，2进制形式的int
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔变量：true 或 false
%c          rune (Unicode码点)，Go语言里特有的Unicode字符类型
%s          string
%q          带双引号的字符串 "abc" 或 带单引号的 rune 'c'
%v          会将任意变量以易读的形式打印出来
%T          打印变量的类型
%%          字符型百分比标志（%符号本身，没有其他操作）


3)
读文件的方式:
bufio.Scanner  
ioutil.ReadFile
bufio.NewReader


package main

import "os"

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/passwd")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}


在更底层一些的地方，bufio.Scanner，ioutil.ReadFile和ioutil.WriteFile使用的都是*os.File的Read和Write方法，不过一般程序员并不需要去直接了解到其底层实现细节，在bufio和io/ioutil包中提供的方法已经足够好用。