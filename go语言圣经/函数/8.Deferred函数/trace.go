package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	fmt.Println("start")
	time.Sleep(10 * time.Second)
	fmt.Println("end")
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
