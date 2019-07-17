package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner()
	fmt.Println(fib(43))
}

func fib(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner() {
	delay := 100 * time.Millisecond
	for {
		for _, c := range "-\\|/" {
			fmt.Printf("\r%c", c)
			time.Sleep(delay)
		}
	}
}
