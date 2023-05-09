package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32 // although should use money package

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// add the handlers
func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("item is already exist: %q", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	db[item] = dollars(p)

	fmt.Fprintf(w, "added %s with price %s\n", item, price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	db[item] = dollars(p)

	fmt.Fprintf(w, "new price %s for item %s\n", price, db[item])
}

func (db database) drop(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	//price := r.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusBadRequest) // 400
		return
	}
	delete(db, item)

	fmt.Fprintf(w, "item %s has been- deleted", item)
}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	// add some routes
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.drop)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
When a function is assigned to a function type, it adopts the method set of that type.
In this case, since http.HandlerFunc is a function type that defines the ServeHTTP method, the assigned MyHandler THAT IS list() function will also have the ServeHTTP method.

Here's a summary of the process:

The MyHandler function has the signature func(http.ResponseWriter, *http.Request).
By assigning MyHandler to http.HandlerFunc, it is implicitly converted to the http.HandlerFunc function type.
The http.HandlerFunc function type includes the ServeHTTP method.
Therefore, the assigned handler variable of type http.HandlerFunc will have the ServeHTTP method, as it adopted the method set of http.HandlerFunc.
This allows the handler to be used as an http.Handler interface, satisfying the contract for handling HTTP requests.
When the server calls the ServeHTTP method on the assigned handler, it will invoke the underlying MyHandler function to handle the request.
*/
