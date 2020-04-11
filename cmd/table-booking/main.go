/*
Copyright 2020 Daniel Ortiz @daniOrtiz11 (https://github.com/daniOrtiz11). All rights reserved.
Code under the MIT License. See LICENSE in the project root for license information.
*/

package main

import (
	"log"
	"net/http"

	"github.com/daniOrtiz11/table-booking/pkg/server"
	"github.com/joho/godotenv"
)

func loadEnv() {
	//change filename to use other environment
	err := godotenv.Load("dev.env")
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
