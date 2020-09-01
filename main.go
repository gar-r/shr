package main

import (
	"log"
)

func main() {
	store, err := NewMongoStore("mongodb://shr_db:27017", "shr")
	if err != nil {
		log.Fatal(err)
	}
	defer store.Disconnect()
}
