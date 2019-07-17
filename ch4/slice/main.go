package main

import "fmt"

func main() {
	s := []string{"a", "b", "c", "d", "e"}
	fmt.Println(s)
	a := s[1:3]
	fmt.Println(a)
	fmt.Println("LEN=", len(a), "CAPACITY=", cap(a))
	b := a[:4]
	fmt.Println(b)

	var another []string
	fmt.Println(another)
	fmt.Println(another == nil)
	fmt.Println(len(another))
	fmt.Println(cap(another))

	// Dynamically allocated slice and append ops
	s2 := make([]string, 0, 2)
	for _, x := range s {
		s2 = append(s2, x)
	}
	fmt.Println(s2, "LEN=", len(s2), "CAPACITY=", cap(s2))

	var x, y []int
	for i := 0; i < 10; i++ {
		y = myAppend(x, i)
		fmt.Println(y, "LEN=", len(y), "CAPACITY=", cap(y))
		x = y
	}

	variadic(10)
	variadic(10, 11)
	variadic(10, 11, 12, 13)
	variadic(10, x...)
}

func myAppend(s []int, n int) []int {
	var z []int
	var newLen = len(s) + 1
	if newLen <= cap(s) {
		z = s[:newLen]
	} else {
		newCap := newLen
		if newCap < 2*len(s) {
			newCap = 2 * len(s)
		}
		if newCap < 4 {
			newCap = 4
		}
		z = make([]int, newLen, newCap)
		copy(z, s)
	}
	z[len(s)] = n
	return z
}

func variadic(x int, y ...int) {
	fmt.Printf("%T %v\n", y, y)
}
