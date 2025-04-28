package main

import (
	"fmt"
	"sync"
)

/* Запуск горутины и ожидание её завершения
Задача: Напишите функцию, которая запускает горутину, выполняющую fmt.Println("Hello from goroutine!"), 
и использует sync.WaitGroup для ожидания её завершения. */

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func ()  {
			fmt.Println("Hello from goroutine!")
			wg.Done()
		}()
		
		wg.Wait()
}

