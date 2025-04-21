package main

import (
	"fmt"
	"sync"
)


func main() {
	wg := &sync.WaitGroup{}
	go func ()  {
			fmt.Println("Hello from goroutine!")
			wg.Done()
		}()
		fmt.Println("Hello from main goroutine!")
		wg.Wait()
}

