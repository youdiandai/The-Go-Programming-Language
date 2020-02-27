package main

import (
	"fmt"
	"os"
	"strings"
)

func expand(s string, f func(string) string) string {
	tmp := strings.Split(s, "")
	for i := range tmp {
		tmp[i] = f(tmp[i])
	}
	return strings.Join(tmp, "")
}
func addStart(s string) string {
	return "*" + s + "*"
}

func main() {
	s := os.Args[1]
	fmt.Println(expand(s, addStart))

}
