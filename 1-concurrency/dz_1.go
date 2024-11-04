package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// начало программы
func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	z := 0
	var array []int
	var array2 []int
	var wg sync.WaitGroup

	//генерация случайных чисел
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			z = rand.Intn(100) + 1
			ch <- z
			array = append(array, z)
		}
		close(ch)
	}()
	//возводим в квадрат
	wg.Add(1)
	go func() {
		defer wg.Done()
		b := 0
		for i := 0; i < 10; i++ {
			b = <-ch
			b = b * b
			array2 = append(array2, b)
			ch2 <- b
		}
		close(ch2)
	}()
	//выводим результат
	for i := 0; i < 10; i++ {
		fmt.Printf("  %d", <-ch2)
	}

}
