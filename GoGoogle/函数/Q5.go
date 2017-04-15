/*
平均值

编写一个函数用于计算一个 float64 类型的 slice 的平均值。
*/

/*
package main

import (
	"fmt"
)

func main() {
	var s []float64
	a := [100]float64{99: 1}
	s = a[1:99]

	sum := 0.0
	for _, v := range s {
		sum += v
	}

	fmt.Println("平均值是", sum/len(s))
}

*/


func average(xs []float64)(avg float64){
	sum :=0.0
	switch len(xs){
	case 0:
		avg=0
	default:
		for _,v :=range xs{
			sum+=v
		}
		avg=sum/float64(len(xs))
	}
	return //命名返回值
}