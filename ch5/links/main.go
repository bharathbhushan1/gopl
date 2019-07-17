package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		fmt.Printf("breadthFirst: processing %d items\n", len(items))
		for _, i := range items {
			if seen[i] {
				continue
			}
			seen[i] = true
			newItems := f(i)
			fmt.Printf("\tAdded %d items\n", len(newItems))
			worklist = append(worklist, newItems...)
		}
	}
}

func crawl(url string) []string {
	fmt.Println("CRAWLING:", url)
	links, err := Extract(url)
	if err != nil {
		log.Printf("crawl error for %s: %s", url, err)
	}
	return links
}

// Extract links from a url
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("got resp status %s for %s", resp.Status, url)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("html parse failed for %s with %v", url, err)
	}
	var links []string
	extractLink := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad urls
				}
				links = append(links, link.String())
			}
		}
	}
	foreachNode(doc, extractLink, nil)
	return links, nil
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
