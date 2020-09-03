package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func router() *mux.Router {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("static"))
	r.HandleFunc("/", ShortenHandler).Methods("POST")
	r.HandleFunc("/{id:[0-9a-zA-Z]+}", DecodeHandler)
	r.Handle("/", fs)
	return r
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	u, err := store.Find(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			handleBadRequest(w, err)
		} else {
			handleInternalServerError(w, err)
		}
		return
	}
	pu, err := parseUrl(u.Val)
	if err != nil {
		handleInternalServerError(w, err)
	} else {
		http.Redirect(w, r, pu, http.StatusSeeOther)
	}
}

func parseUrl(s string) (string, error) {
	tu, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	if tu.Scheme == "" {
		tu.Scheme = "http"
	}
	return tu.String(), nil
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if param, err := getRequestData(r); err == nil {
		url, err := store.FindByUrl(param)
		if err == nil {
			writeResponse(w, r, url.Id)
			return
		} else if errors.Is(err, mongo.ErrNoDocuments) {
			if sha, err := Shorten(param); err == nil {
				if err = store.Save(&Url{Id: sha, Val: param}); err == nil {
					writeResponse(w, r, sha)
					return
				}
			}
		}
	}
	handleInternalServerError(w, err)
}

func getRequestData(r *http.Request) (s string, err error) {
	body := new(bytes.Buffer)
	_, err = body.ReadFrom(r.Body)
	s = body.String()
	return
}

func writeResponse(w http.ResponseWriter, r *http.Request, sha string) {
	res := fmt.Sprintf("%s%s%s", r.Host, r.URL.Path, sha)
	if _, err := w.Write([]byte(res)); err != nil {
		handleInternalServerError(w, err)
	}
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
