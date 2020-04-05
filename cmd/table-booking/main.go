package main

import (
	"log"
	"net/http"

	"github.com/daniOrtiz11/table-booking/pkg/server"
	"github.com/joho/godotenv"
)

/*
var c = make(chan int)
var quit = make(chan int)
*/

func loadEnv() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error loading env file: ", err)
	}
}

func main() {
	loadEnv()
	log.Println("Starting server...")
	s := server.New()
	log.Fatal(http.ListenAndServe(s.Addr(), s.Router()))
}
