package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordsCounter int
type LinesCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		*c++
	}
	return len(p), nil
}

func (c *LinesCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanLines)
	for input.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	var wc WordsCounter
	wc.Write([]byte("hello world"))
	fmt.Println(wc)
	var lc LinesCounter
	lc.Write([]byte("hello world\ngo golang\ntoday\nhappy"))
	fmt.Println(lc)
}
