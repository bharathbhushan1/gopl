package main

import (
	"fmt"
	"sort"
)

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}
func (s stringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := stringSlice{"b", "c", "a"}
	sort.Sort(s)
	fmt.Println(s)
}
