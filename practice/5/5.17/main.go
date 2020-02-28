package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagname(n *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node
	var elementsByTagname func(*html.Node, string)
	elementsByTagname = func(n *html.Node, name string) {
		if n.Type == html.ElementNode && n.Data == name {
			fmt.Printf("finded:%s", name)
			nodes = append(nodes, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			elementsByTagname(c, name)
		}
	}
	for _, v := range names {
		fmt.Printf("search %s\n", v)
		elementsByTagname(n, v)
	}
	return nodes
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Elements by tag name:%v", err)
		os.Exit(1)
	}
	nodes := ElementsByTagname(doc, "main", "meta")
	for _, n := range nodes {
		fmt.Printf("%s\n", n.Data)
	}
}
