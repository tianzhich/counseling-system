package main

import (
	"counseling-system/pkg/feature1"
	"counseling-system/pkg/oauth"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/fake", feature1.Handler)

	oauthHandler(mux)

	log.Println("Listening on port 8081 ...")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}

func oauthHandler(mux *http.ServeMux) {
	mux.HandleFunc("/api/signup", oauth.SignupHandler)
}
