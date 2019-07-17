package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	numbers := [...]string{"", "1", "12", "123", "1234", "12345", "123456", "1234567", "12345678901234567890"}
	for _, n := range numbers {
		fmt.Println(comma2(n))
	}
	fmt.Println(isAnagram("abcd", "bdca"))
}

func comma(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}
	return comma(s[:l-3]) + "," + s[l-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer
	l := len(s)
	if l <= 3 {
		return s
	}
	var i int
	if (l-3)%2 == 0 {
		i = 0
	} else {
		i = 1
		buf.WriteByte(s[0])
		buf.WriteRune(',')
	}
	for ; i < (l - 3); i += 2 {
		buf.WriteString(s[i : i+2])
		buf.WriteRune(',')
	}
	// Write last 3 digits
	buf.WriteString(s[l-3:])
	return buf.String()
}

func isAnagram(a string, b string) bool {
	s := strings.Split(a, "")
	sort.Strings(s)
	a = strings.Join(s, "")

	s = strings.Split(b, "")
	sort.Strings(s)
	b = strings.Join(s, "")
	return a == b
}
