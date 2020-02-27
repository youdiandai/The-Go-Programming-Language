package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var flag bool = false
var target *html.Node

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, m string) bool) {
	fmt.Println("exce") //用来测试找到后是否尽快结束了递归
	if !flag {          //如果不加这一层循环，在找到后还会对兄弟节点进行多次循环
		if pre != nil {
			flag = pre(n, id)
		}
		if !flag {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				forEachNode(c, id, pre, post)
			}
		} else {
			target = n
		}

		if post != nil {
			post(n, id)
		}
	}

}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, k := range n.Attr {
			if k.Key == "id" && k.Val == id {
				fmt.Println("finded")
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, k := range n.Attr {
			if k.Key == "id" && k.Val == id {
				return true
			}
		}
	}
	return false
}

func ElementByID(doc *html.Node, id string) *html.Node {
	forEachNode(doc, id, startElement, endElement)
	return target
}

func main() {
	id := os.Args[1]
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "element by id:%v", err)
		os.Exit(1)
	}
	n := ElementByID(doc, id)
	if n != nil {
		for _, a := range n.Attr {
			if a.Key == "id" {
				fmt.Printf("element name:%s id==%s", n.Data, id)
			}
		}
	} else {
		fmt.Println("not find")
	}
}

//完成的还不够好，应该还有更好的方法
