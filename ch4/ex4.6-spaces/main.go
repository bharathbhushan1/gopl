package main

import "fmt"

func main() {
	fmt.Printf("%q\n", squashSpaces("hello   world   "))
}

func squashSpaces(s string) string {
	return s
}
