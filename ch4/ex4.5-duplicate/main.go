package main

import "fmt"

func main() {
	s := []int{1, 1, 1, 1, 2, 2, 3, 3, 4, 5, 6, 5, 4}
	fmt.Println(removeDuplicates2(s))
}

func removeDuplicates(s []int) []int {
	if len(s) == 0 {
		return s
	}
	var result []int
	cur := s[0]
	result = append(result, cur)
	for i := 1; i < len(s); i++ {
		if s[i] == cur {
			continue
		}
		cur = s[i]
		result = append(result, cur)
	}
	return result
}

func removeDuplicates2(s []int) []int {
	num := 0
	var cur int
	for i := 0; i < len(s); i++ {
		if i == 0 || cur != s[i] {
			s[num] = s[i]
			cur = s[num]
			num++
		}
	}
	return s[:num]
}
