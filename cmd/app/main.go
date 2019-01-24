package main

import (
	"counseling-system/pkg/feature1"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type res struct {
	Data string `json:"data"`
	Code int    `json:"code"`
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Hello World, %s! You are %s, and %d years old", feature1.GetPerson().Name, feature1.GetPerson().Gender, feature1.GetPerson().Age)
	var res = res{s, 200}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "JSON")
	result, _ := json.Marshal(res)

	fmt.Fprintf(w, string(result))

	db, err := sql.Open("mysql", "tianzhi:tianzhi@tcp(47.94.223.143:3306)/pcs")
	checkErr(err)

	rows, err := db.Query("select username, password, phone from user")
	checkErr(err)

	for rows.Next() {
		var username string
		var password string
		var phone string

		err = rows.Scan(&username, &password, &phone)
		checkErr(err)
		fmt.Println(username, password, phone)
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", personHandler)

	log.Println("Listening on port 8080 ...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
