package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotateLeft(s, 2)
	fmt.Println(s)
}

func rotateLeft(s []int, n int) {
	if n == 0 {
		return
	}
	n = n % len(s)
	temp := make([]int, n, n)
	copy(temp, s[:n])
	copy(s[:len(s)-n], s[n:])
	copy(s[len(s)-n:], temp)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
