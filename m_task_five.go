package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var increment int32

func inc(increment *int32) {
	atomic.AddInt32(increment, 1)
}

func main() {
	
	for i := 0; i < 10; i++ {
		go inc(&increment)
	}
	time.Sleep(time.Second)
	fmt.Println(increment)
}

