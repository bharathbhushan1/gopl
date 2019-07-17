package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "\U00000CA0-\U00000CA0"
	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Printf("% x\n", str)
	fmt.Printf("% x\n", []rune(str))
	fmt.Println(basename("/a/b/c.go"))
	fmt.Println(basename(""))
	fmt.Printf("%q\n", basename(".    /    ."))
	fmt.Printf("%q\n", basename("/"))
	fmt.Println(basename2("/a/b/c.go"))
	fmt.Println(basename2(""))
	fmt.Printf("%q\n", basename2(".    /    ."))
	fmt.Printf("%q\n", basename2("/"))
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	i := strings.LastIndex(s, "/")
	s = s[i+1:]
	i = strings.LastIndex(s, ".")
	if i >= 0 {
		s = s[:i]
	}
	return s
}
