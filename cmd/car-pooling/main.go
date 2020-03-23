package main

import (
	"log"
	"net/http"

	"github.com/daniOrtiz11/car-pooling/pkg/server"
	"github.com/joho/godotenv"
)

/*
var c = make(chan int)
var quit = make(chan int)
*/

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading env file")
	}
}

func main() {
	loadEnv()
	log.Println("Starting server...")
	s := server.New()
	log.Fatal(http.ListenAndServe(s.Addr(), s.Router()))
}
