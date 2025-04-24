package main

import (
	"fmt"
	"sync"
)


func printNumber(wg *sync.WaitGroup, ch chan<- int ) {

	defer wg.Done()

	for i := 1; i <=5 ; i++ {
		ch <-i
	}
	close(ch)
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(1)
 
	go printNumber(&wg, ch)

	sum := 0
	for i := range ch {
		sum += i
	}
	fmt.Println(sum)
	wg.Wait()

}

