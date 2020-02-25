package main

/**
- 4.8：修改 charcount 的代码来统计字母、数字和其他在 Unicode 分类中的字符数量，可以使用 unicode.IsLetter 等
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[string]int) // Unicode 字符数量
	counts["letter"] = 0
	counts["num"] = 0
	counts["others"] = 0
	counts["invalid"] = 0
	var utflen [utf8.UTFMax + 1]int // UTF-8 编码的长度

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			counts["invalid"]++
			continue
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		}
		if unicode.IsNumber(r) {
			counts["num"]++
		}
		counts["others"]++
		utflen[n]++
	}
	fmt.Printf("type\t\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t\t%d\n", c, n)
	}
	fmt.Printf("\nlrn\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
}
