package main

import (
	"log"
	"net/http"
	"regexp"
)

func router() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.HandleFunc("/s", shortenHandler)
	mux.HandleFunc("/r/", redirectHandler)
	mux.Handle("/", fs)
	return mux
}

var pattern = regexp.MustCompile("/r/(.+)")

func redirectHandler(w http.ResponseWriter, r *http.Request) {

	m := pattern.FindStringSubmatch(r.URL.Path)
	if m == nil || len(m) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sha := m[1]
	url, err := decode(sha)
	if err != nil {
		if isNotFound(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := readBody(r)
	if err != nil || body == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := parseUrl(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sha, err := encode(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	_, err = w.Write([]byte(fqShortUrl(r, sha)))
	if err != nil {
		log.Println(err)
	}

}
