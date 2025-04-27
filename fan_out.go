package main

import (
	"fmt"
	"sync"
)

/* 
Fan-Out (разделение работы между несколькими воркерами)
Задача: Создайте функцию, которая принимает канал с задачами и распределяет их между N горутинами.
// Разделить канал на n каналов,
// которые получают сообщения в циклическом порядке.
func Split(ch <-chan int, n int) []<-chan int
*/


func Split(ch <-chan int, n int) []<-chan int {
	cs := make([]chan int, n)
	for i := 0; i < n; i++ {
	  cs[i] = make(chan int)
	}

	go func() {
		defer func() {
			for _, c := range cs {
				close(c)
			}
		}()

		i := 0
		for v := range ch {
			cs[i] <- v
			i = (i + 1) % n
		}
	}()
	
	out := make([]<-chan int, n)
	for i , c := range cs {
		out[i] = c
	}
	return out
}




func main() {
	input := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	outputs := Split(input, 3)

	var wg sync.WaitGroup
	for i, ch  := range outputs {
		wg.Add(1)
		go func(i int, ch <-chan int) {
			defer wg.Done()
			for v:= range ch {
				fmt.Printf("Воркер %d получен: %d\n", i, v)
			}
		}(i, ch)
	}
	wg.Wait()
	
}
