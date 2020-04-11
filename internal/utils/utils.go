package utils

import (
	"net/http"
	"os"
)

/*
WAITING ENUMERATED STATUS
*/
const WAITING = 1

/*
EATING ENUMERATED STATUS
*/
const EATING = 2

/*
COMPLETED ENUMERATED STATUS
*/
const COMPLETED = 3

/*
GetContentType will retrieve the content type header after successful search in request.
In other case, will retrieve a default content type.
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
GetAccept will retrieve the accept header after successful search in request.
In other case, will retrieve a default accept.
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
GetEnv will retrieve the value by key after successful search in env file.
In other case, will retrieve the fallback value.
*/
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
