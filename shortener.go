package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
)

func Shorten(url string) (s string, err error) {
	for l := 2; l <= 20; l++ {
		if s, err = shorten(url, l); store.Missing(s) || err != nil {
			return
		}
	}
	err = errors.New("cannot hash url")
	return
}

func shorten(url string, len int) (string, error) {
	hash := sha1.New()
	if _, err := io.WriteString(hash, url); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)[:len]), nil
}
