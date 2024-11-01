package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	z := 0
	var array []int
	var array2 []int
	go func() {

		for i := 0; i < 10; i++ {
			z = rand.Intn(100) + 1
			ch <- z
			array = append(array, z)
		}
	}()

	go func() {
		b := 0
		for i := 0; i < 10; i++ {
			b = <-ch
			b = b * b
			array2 = append(array2, b)
			ch2 <- b
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("  %d", <-ch2)
	}

}
