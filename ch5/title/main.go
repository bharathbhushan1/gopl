package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func main() {
	for _, k := range os.Args[1:] {
		err := title(k)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

func trace(name string) func() {
	start := time.Now()
	log.Printf("START: %s", name)
	return func() { log.Printf("END %s (%vs)\n", name, time.Since(start).Seconds()) }
}

func title(url string) error {
	defer trace("title")()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http response status for %s is %s", url, resp.Status)
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s response is not text/html but %s", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("html parsing failed for %s: %s", url, err)
	}
	result := findTitle(doc)
	fmt.Println(result)
	return nil
}

func findTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title := findTitle(c)
		if title != "" {
			return title
		}
	}
	return ""
}
