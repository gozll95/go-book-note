#局部变量仅仅在执行定义它的函数时有效


/*
var a int 

func main(){
	a=5
	println(a)
	f()
}

func f(){
	a:=6
	println(a)
	g()
}

func g(){
	println(a)
}

*/

\\
#defer延迟函数是按照后进先出(LIFO)的顺序执行
for i:=0;i<5;i++{
	defer fmt.Printf("%d",i)
}



#defer的作用:
1)延迟
2)修改返回值

defer func(){
	...
}

defer func(x int){
	/*...*/
}(5)



\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

func f()(ret int){
	defer func(){
		ret++   //ret增加为1
	}()
	return 0 //返回的是1而不是0
}

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\


#变参
func myfunc(arg ...int){
	for _,n:=range arg{
		fmt.Printf("And the number is: %d\n",n)
	}
}

变参是slice,如果不指定变参的类型是空接口。

 func myfunc(arg ...int){
 	myfunc2(arg...)
 	myfunc2(arg[:2]...)
 }



 #panic-recover

 func throwsPanic(f func())(b bool){
 	defer func(){
 		if x:=recover();x!=nil{
 			b=true
 		}
 	}()

 	f()

 	return
 }

 