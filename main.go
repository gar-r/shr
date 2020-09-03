package main

import (
	"log"
	"net/http"
	"os"
)

const addr = ":17800"
const dir = "urls"

func main() {
	mkdir()
	serve()
}

func mkdir() {
	if err := os.Mkdir(dir, 0744); err != nil {
		if !os.IsExist(err) {
			log.Fatal(err)
		}
	}
}

func serve() {
	log.Fatal(http.ListenAndServe(addr, router()))
}
