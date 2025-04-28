package main

import (
	"fmt"
	"sync"
)

/* Запуск нескольких горутин и ожидание их завершения
Задача: Напишите программу, которая запускает 5 горутин, 
каждая из которых печатает свой номер (от 1 до 5), 
и использует sync.WaitGroup для их синхронизации(нужно подождать их выполнения). */

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