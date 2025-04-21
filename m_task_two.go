package main

import (
	"fmt"
	"sync"
)


func printNumber(wg *sync.WaitGroup) {
//	defer wg.Done()
	for i := 1; i<=5 ; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(5)
//	for i := 0; i < 5; i++ {
//		go printNumber(wg)
//	}
	go printNumber(wg)
	go printNumber(wg)
	go printNumber(wg)
	go printNumber(wg)
	go printNumber(wg)
	wg.Wait()
}

