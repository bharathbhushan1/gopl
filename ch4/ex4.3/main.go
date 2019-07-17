package main

import "fmt"

func main() {
	s := [5]int{1, 2, 3, 4, 5}
	reverse(&s)
	fmt.Println(s)
}

func reverse(p *[5]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		p[i], p[j] = p[j], p[i]
	}
}
