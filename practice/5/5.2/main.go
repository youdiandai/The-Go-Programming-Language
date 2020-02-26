/**
- 5.2：写一个函数，用于统计HTML文档树内所有元素个数，如p,div,span等
*/
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "element count:%v\n", err)
		os.Exit(1)
	}
	counter := make(map[string]int)
	elementCount(counter, doc)
	fmt.Printf("element_type\t\tnum\n")
	for k, v := range counter {
		fmt.Printf("%s\t\t%d\n", k, v)
	}

}

func elementCount(counter map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elementCount(counter, c)
	}
}
