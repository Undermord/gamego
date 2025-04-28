package main

import (
	"fmt"
	"sync"
)

/* 
Fan-In (объединение данных из нескольких каналов)
Задача: Напишите функцию, которая объединяет два входных канала в один выходной.
// Объединяем разные каналы в один канал
func Merge(cs ...<-chan int) <-chan int
*/


func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out :=make (chan int)

	send := func(c <- chan int){
		for n := range c {
			out <-n
		}
		wg.Done()
	}
	wg.Add(len(cs))

	for _, c := range cs {
		go send(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}



func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i <= 5; i++ {
			ch1 <- i
		}
		close(ch1)

	}()
	go func() {
		for i := 5; i <= 10; i++ {
		ch2 <- i
		}
		close(ch2)
	}()

	merged := Merge (ch1,ch2)
	for n := range merged {
		fmt.Println(n)
	}
	
}
