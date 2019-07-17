package main

import (
	"fmt"
	"log"
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
	prereqs["computer organization"] = []string{"compilers"}
	fmt.Println(prereqs)
	for i, course := range toposort(prereqs) {
		fmt.Printf("%d\t%s\n", i+1, course)
	}
}

func toposort(m map[string][]string) []string {
	var result []string
	seen := make(map[string]bool)
	active := make(map[string]bool)
	depth := 0

	var visitAll func([]string)
	visitAll = func(items []string) {
		depth++
		for _, k := range items {
			fmt.Printf("%*s STARTING %s\n", depth*2, "", k)
			if active[k] {
				log.Fatalf("cycle detected at %s\n", k)
			}
			if !seen[k] {
				seen[k] = true
				active[k] = true
				fmt.Printf("%*s VISITING CHILDREN OF %s\n", depth*2, "", k)
				visitAll(m[k])
				active[k] = false
				result = append(result, k)
			}
		}
		depth--
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	visitAll(keys)
	return result
}
