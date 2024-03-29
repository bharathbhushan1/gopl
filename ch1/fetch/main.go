// Fetch and print the content at a URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "GET: url=%s, error=%v\n", url, err)
			continue
		}

		/*
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "READ: url=%s, error=%v\n", url, err)
				continue
			}
			fmt.Printf("%s", b)
		*/

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "COPY: url=%s, error=%v\n", url, err)
			continue
		}
	}
}
