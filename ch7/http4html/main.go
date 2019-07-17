package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var dbMutex sync.Mutex

const listHTMLTemplateStr = `
<style>
table {border: 1px solid black;}
th {border: 1px solid black; padding: 5px; text-align: left; background-color: #4CAF50; color: white;}
td {border: 1px solid black; padding: 5px; text-align: left;}
</style>
<table>
	<tr><th>Item</th><th>Price</th></tr>
	{{range $k, $v := .}}
	<tr><td>{{$k}}</td><td>{{$v}}</td>
	{{end}}
</table>
`

var listHTMLTemplate = template.Must(
	template.New("listHtml").Parse(listHTMLTemplateStr))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	dbMutex.Lock()
	listHTMLTemplate.Execute(w, db)
	dbMutex.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	i := req.URL.Query().Get("item")
	dbMutex.Lock()
	price, ok := db[i]
	dbMutex.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found %q", i)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", i, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	i := req.URL.Query().Get("item")
	newPriceStr := req.URL.Query().Get("price")
	newPrice, err := strconv.Atoi(newPriceStr)

	dbMutex.Lock()
	defer dbMutex.Unlock()
	price, ok := db[i]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found %q", i)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "value %q invalid, not updating price for %s", newPriceStr, i)
		return
	}
	fmt.Fprintf(w, "Updated price of %s from %s to %s\n", i, price, dollars(newPrice))
	db[i] = dollars(newPrice)
}
