package main

import (
	"fmt"
	"sync"
)



func incrementValue( increment *int,wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	
	mu.Lock()
	*increment++
	fmt.Println(*increment)
	mu.Unlock()

}


func main() {
	increment := 1
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementValue(&increment, wg, mu)

	}
	wg.Wait()
}

