package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{id}", DecodeHandler)
	r.HandleFunc("/", ShortenHandler).Methods("POST")
	r.HandleFunc("/", RootHandler).Methods("GET")
	return r
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url, err := store.Find(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			handleBadRequest(w, err)
		}
		handleInternalServerError(w, err)
		return
	}
	http.Redirect(w, r, url.Val, http.StatusFound)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		handleBadRequest(w, err)
		return
	}
	url := r.Form["url"][0]
	s, err := Shorten(url)
	if err != nil {
		handleInternalServerError(w, err)
		return
	}
	if err = store.Save(&Url{Id: s, Val: url}); err != nil {
		handleInternalServerError(w, err)
	}
	if _, err = w.Write([]byte(s)); err != nil {
		handleInternalServerError(w, err)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("TODO")); err != nil {
		handleInternalServerError(w, err)
	}
	w.WriteHeader(http.StatusOK)
}

func handleBadRequest(w http.ResponseWriter, err error) {
	handleError(w, err, http.StatusBadRequest)
}

func handleInternalServerError(w http.ResponseWriter, err error) {
	handleError(w, err, http.StatusInternalServerError)
}

func handleError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	log.Printf("%v\n", err)
}
