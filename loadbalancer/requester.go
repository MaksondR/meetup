package main

import (
	"fmt"
	"time"
)

func doSomeWork() int {
	n := time.Duration(int64(time.Second))
	time.Sleep(n)

	fmt.Println("Sleep")

	return int(n)
}

type Request struct {
	fn func() int
	c chan int
}

func requester(work chan Request) {
	c := make(chan int)
	for {
		work <- Request{
			fn: doSomeWork,
			c:  c,
		}
		<-c
	}
}
