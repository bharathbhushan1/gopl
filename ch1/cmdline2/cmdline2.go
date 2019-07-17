// This example prints command line args using range function
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg + "(" + fmt.Sprintf("%d", i) + ")"
		sep = " "
	}
	fmt.Println(s)

	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], "-"))
	fmt.Println(os.Args[1:])
}
