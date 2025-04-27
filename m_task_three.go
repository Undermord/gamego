package main

import (
	"fmt"
	"sync"
)
/* Использование канала для передачи данных
Задача: Напишите функцию, которая создает горутину, отправляющую числа от 1 до 5 в канал,
 а затем в main извлекает их и складывает, результат выводит в консоль. */

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

