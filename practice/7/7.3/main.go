package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	writeBuf(&buf, t)
	str := buf.String()
	str = str[:len(str)-1] + "]"
	return str
}

func writeBuf(buf *bytes.Buffer, t *tree) {
	if t != nil {
		writeBuf(buf, t.left)
		fmt.Fprintf(buf, "%d,", t.value)
		writeBuf(buf, t.right)
	}
}

func sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value <= t.value {
		t.left = add(t.left, value)
	}
	if value > t.value {
		t.right = add(t.right, value)
	}
	return t
}
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func main() {
	test := []int{3, 4, 5, 6, 7, 3, 4, 5, 45, 446, 75, 43, 321, 46, 4, 22, 76}
	var testTree *tree
	for _, v := range test {
		testTree = add(testTree, v)
	}
	fmt.Println(testTree)
}
