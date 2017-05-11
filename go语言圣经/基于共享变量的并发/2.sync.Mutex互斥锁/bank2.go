package main

var (
	sema    = make(chan struct{}, 1) // a binary seamphore guarding balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} //acquire token
	balance = balance + amount
	<-sema //release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}
