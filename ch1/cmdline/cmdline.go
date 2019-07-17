// This example prints command line args
// Uses os package
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// os.Args is a slice
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
