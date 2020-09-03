package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func readBody(r *http.Request) (s string, err error) {
	body := new(bytes.Buffer)
	if _, err = body.ReadFrom(r.Body); err != nil {
		return
	}
	return body.String(), nil
}

func parseUrl(s string) (string, error) {
	parsed, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	if parsed.Scheme == "" {
		parsed.Scheme = "http"
	}
	return parsed.String(), nil
}

func fqShortUrl(r *http.Request, sha string) string {
	return fmt.Sprintf("%s%s/%s", r.Host, "r", sha)
}
