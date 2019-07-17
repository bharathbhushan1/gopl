package main

import (
	"fmt"
	"sort"
)

func main() {
	m1 := map[string]int{}
	fmt.Println(m1)

	m2 := map[string]int{
		"mesha":     10,
		"vrushabha": 20,
	}
	fmt.Println(m2)
	m2["mithuna"] = 30
	m2["karkataka"] = 40
	fmt.Println(m2)

	fmt.Println("--- ITERATION")
	for k, v := range m2 {
		fmt.Println("KEY=", k, "VALUE=", v)
	}

	var keys []string
	for k := range m2 {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("--- SORTING")
	for _, k := range keys {
		fmt.Println("KEY=", k, "VALUE=", m2[k])
	}

	fmt.Println("--- DELETION")
	delete(m2, "mithuna")
	fmt.Println(m2)
}
