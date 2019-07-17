// Find duplicate lines
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		contents, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue // to next file
		}
		for _, line := range strings.Split(string(contents), "\n") {
			if len(line) > 0 {
				counts[line]++
			}
		}
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
