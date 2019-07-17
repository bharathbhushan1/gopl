package main

import "fmt"

func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 10; i++ {
		out <- i
	}
}

func squarer(out chan<- int, in <-chan int) {
	defer close(out)
	for x := range in {
		out <- x * x
	}
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Printf("%d ", x)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
