package main

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	NotFound = errors.New("not found")
	NoDigits = errors.New("no digits left")
)

func encode(url string) (string, error) {
	sha, err := hash(url)
	if err != nil {
		return "", err
	}
	for l := 4; l <= len(sha); l++ {
		frag := sha[:l]
		name := fname(frag)
		if exists(name) {
			b, err := ioutil.ReadFile(name)
			if err != nil {
				return "", err
			}
			if string(b) == url {
				return frag, nil
			}
			continue
		}
		err = ioutil.WriteFile(name, []byte(url), 0644)
		if err != nil {
			return "", err
		}
		return frag, nil
	}
	return "", NoDigits
}

func decode(id string) (string, error) {
	name := fname(id)
	if !exists(name) {
		return "", NotFound
	}
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func hash(url string) (string, error) {
	hash := sha1.New()
	if _, err := io.WriteString(hash, url); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func fname(s string) string {
	return fmt.Sprintf("%s/%s", dir, s)
}

func exists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func isNotFound(err error) bool {
	return err.Error() == NotFound.Error()
}
