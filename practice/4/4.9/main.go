package main

/**
- 4.9：编写一个程序wordfreq来汇总输入文本文件中每个单词出现的次数。在第一次调用Scan之前，需要使用input.Split(bufio.ScanWords)来讲文本行按照单词分割而不是行分割
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		seen[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq :%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("word\tnum\n")
	for k, v := range seen {
		fmt.Printf("%s\t%d\n", k, v)
	}

}
