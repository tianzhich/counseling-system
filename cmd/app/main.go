package main

import (
	"counseling-system/pkg/feature1"
	"counseling-system/pkg/register"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/fake", feature1.PersonHandler)
	mux.HandleFunc("/api/register", register.RegisterHandler)

	log.Println("Listening on port 8080 ...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
