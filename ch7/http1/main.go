package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for i, p := range db {
			fmt.Fprintf(w, "%s: %s\n", i, p)
		}
	case "/price":
		i := req.URL.Query().Get("item")
		price, ok := db[i]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "item not found %q", i)
			return
		}
		fmt.Fprintf(w, "%s: %s\n", i, price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page %s", req.URL)
	}
}
