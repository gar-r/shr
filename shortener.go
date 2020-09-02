package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"

	"go.mongodb.org/mongo-driver/mongo"
)

func Shorten(url string) (s string, err error) {
	for l := 2; l <= 20; l++ {
		if s, err = shorten(url, l); err != nil {
			return
		}
		var ex bool
		if ex, err = exists(s); err != nil || !ex {
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

func exists(id string) (bool, error) {
	_, err := store.Find(id)
	return !errors.Is(err, mongo.ErrNoDocuments), err
}
