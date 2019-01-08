package main

import (
	"counseling-system/pkg/feature1"
	"fmt"
	"log"
	"net/http"
)

func personHandler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Hello World, %s! You are %s, and %d years old", feature1.GetPerson().Name, feature1.GetPerson().Gender, feature1.GetPerson().Age)
	fmt.Fprintf(w, s)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", personHandler)

	log.Println("Listening on port 8080 ...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
