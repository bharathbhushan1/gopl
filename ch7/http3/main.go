package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	server := http.NewServeMux()
	server.HandleFunc("/list", db.list)
	server.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", server))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for i, p := range db {
		fmt.Fprintf(w, "%s: %s\n", i, p)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	i := req.URL.Query().Get("item")
	price, ok := db[i]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found %q", i)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", i, price)
}
