package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	var db = database{"shoes" : 50, "socks" : 5}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

type dollars float32

func (d dollars) String() string{
	return fmt.Sprintf("%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request){
	for item, price := range db{
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request){
	stuff := req.URL.Query().Get("cc")
	price, ok := db[stuff]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", stuff)
		return
	}
	fmt.Fprintf(w, "%s\n", price)


}
