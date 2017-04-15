package main

import "fmt"

type S struct {
	i int
}

func (p S) Get() int {
	return p.i
}

func (p S) Put(v int) {
	p.i = v
}

type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	fmt.Println(p.Get())
	p.Put(1)
}

func g(something interface{}) int {
	return something.(I).Get()
}

func main() {
	var p I
	p = S{i: 1}
	f(p)

	s := new(S)
	fmt.Println(g(s))
}
