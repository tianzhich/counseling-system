package main

import (
	"counseling-system/pkg/info"
	"counseling-system/pkg/oauth"

	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// API handler
	oauthHandler(mux)
	infoHandler(mux)

	log.Println("Listening on port 8081 ...")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}

func oauthHandler(mux *http.ServeMux) {
	mux.HandleFunc("/api/oauth/signup", oauth.SignupHandler)
	mux.HandleFunc("/api/oauth/signin", oauth.SigninHandler)
	mux.HandleFunc("/api/oauth/auth", oauth.AuthHandler)
	mux.HandleFunc("/api/oauth/signout", oauth.SignoutHandler)
	mux.HandleFunc("/api/oauth/apply", oauth.ApplyHandler)
}

func infoHandler(mux *http.ServeMux) {
	mux.HandleFunc("/api/info/counselingFilters", info.CounselorFilterHandler)
}
