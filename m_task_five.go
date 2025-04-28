package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/* Потокобезопасный инкремент - Atomic.
Задача: Напишите программу, где 10 горутин инкрементируют один счётчик без использования мютексов,
 через атомики. */

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

