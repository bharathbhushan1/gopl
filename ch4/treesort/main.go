package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := []int{5, 6, 4, 2, 1, 3, 8, -1, 9, 4}
	fmt.Println(s)
	s = sort(s)
	fmt.Println(s)
}

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) printTree(buf *bytes.Buffer) {
	if t != nil {
		t.left.printTree(buf)
		strval := fmt.Sprintf("%d ", t.value)
		buf.WriteString(strval)
		t.right.printTree(buf)
	}
}

func (t *tree) String() string {
	var buf bytes.Buffer
	t.printTree(&buf)
	return buf.String()
}

// Add an element into the tree at the right place BST
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(s []int, t *tree) []int {
	if t != nil {
		s = appendValues(s, t.left)
		s = append(s, t.value)
		s = appendValues(s, t.right)
	}
	return s
}

func sort(s []int) []int {
	var root *tree
	for _, v := range s {
		root = add(root, v)
	}
	appendValues(s[:0], root)
	fmt.Println(root)
	return s
}
