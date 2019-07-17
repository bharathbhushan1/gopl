package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

var depth = 0

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	outline(nil, doc)

	var counts = make(map[string]int)
	countElements(counts, doc)
	for k, v := range counts {
		fmt.Printf("%-10.10s %3d\n", k, v)
	}

	depth = 0
	foreachNode(doc, startElement, endElement)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func countElements(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countElements(counts, c)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func foreachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		foreachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
