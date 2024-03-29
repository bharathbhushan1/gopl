package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range toposort(prereqs) {
		fmt.Printf("%d\t%s\n", i, course)
	}
}

func toposort(m map[string][]string) []string {
	var result []string
	seen := make(map[string]bool)

	var visitAll func([]string)
	visitAll = func(items []string) {
		for _, k := range items {
			if !seen[k] {
				seen[k] = true
				visitAll(m[k])
				result = append(result, k)
			}
		}
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	visitAll(keys)
	return result
}
