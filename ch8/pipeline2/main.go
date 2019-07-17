package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		defer close(naturals)
		for i := 0; i < 10; i++ {
			naturals <- i
		}
	}()

	go func() {
		defer close(squares)
		for x := range naturals {
			squares <- x * x
		}
	}()

	for y := range squares {
		fmt.Println(y)
	}
}
