/**
- 5.3：写一个函数，用于输出HTML文档树中所有文本节点的内容。但不包括<script>或<style>元素，因为这些内容在web浏览器中是不可见的。
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
		fmt.Fprintf(os.Stderr, "content print :%v\n", err)
	}
	contentPrint(doc)
}

func contentPrint(n *html.Node) {
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" {
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode && n.FirstChild.NextSibling == nil {
			fmt.Printf("content:%s\n", n.FirstChild.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		contentPrint(c)
	}
}
