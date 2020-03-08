package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getContentType(r *http.Request) string {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	return contentType
}

func unMarshalGroup(b []byte) (group, error) {
	var jsonBlob = []byte(`
	{"Name": "Platypus", "Order": "Monotremata"}
`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}

	var g group
	var jsonBlob2 = []byte(`{
		"id": 1,
		"people": 4,
		"journeying": true
	  }`)
	if err := json.Unmarshal(jsonBlob2, &g); err != nil {
		return g, err
	}
	g.journeying = false
	return g, nil
}
