package main

import (
	"counseling-system/pkg/feature1"
	"counseling-system/pkg/register"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/fake", feature1.Handler)
	mux.HandleFunc("/api/register", register.Handler)

	log.Println("Listening on port 8081 ...")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
