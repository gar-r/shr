package main

import (
	"log"
	"net/http"
)

const connStr = "mongodb://shr_db:27017"
const db = "shr"

const addr = ":17800"

var store *Store

func main() {
	store = initStore()
	defer store.Disconnect()
	r := router()
	log.Fatal(http.ListenAndServe(addr, r))
}

func initStore() *Store {
	store, err := NewMongoStore(connStr, db)
	if err != nil {
		log.Fatal(err)
	}
	return store
}
