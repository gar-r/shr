package main

import "crypto/md5"

func Shorten(url string, len int) string {
	hash := md5.Sum([]byte(url))
	return string(hash[:len])
}
