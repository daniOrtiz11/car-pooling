package utils

import (
	"net/http"
	"os"
)

/*
WAITING is a
*/
const WAITING = 1

/*
EATING is a
*/
const EATING = 2

/*
COMPLETED is a
*/
const COMPLETED = 3

/*
GetContentType is a
*/
func GetContentType(r *http.Request) string {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		//set default value
		contentType = "application/octet-stream"
	}
	return contentType
}

/*
GetAccept is a
*/
func GetAccept(r *http.Request) string {
	contentType := r.Header.Get("Accept")
	if contentType == "" {
		//set default value
		contentType = "application/octet-stream"
	}
	return contentType
}

/*
GetEnv is a
*/
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
