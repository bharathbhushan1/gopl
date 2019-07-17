package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	reverse(a)
	fmt.Println(a)
	reverse(a)
	fmt.Println(a)
	rotateLeft(a, 0)
	fmt.Println(a)
	rotateLeft(a, 1)
	fmt.Println(a)
	rotateLeft(a, 2)
	fmt.Println(a)
	rotateRight(a, 2)
	fmt.Println(a)
	rotateRight(a, 1)
	fmt.Println(a)
	rotateRight(a, 0)
	fmt.Println(a)
}

// In-place reversal of a slice of integers
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateLeft(s []int, n int) {
	fmt.Println("LEFT ROTATING BY", n)
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rotateRight(s []int, n int) {
	fmt.Println("RIGHT ROTATING BY", n)
	reverse(s)
	reverse(s[:n])
	reverse(s[n:])
}
