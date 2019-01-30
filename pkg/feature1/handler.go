package feature1

import (
	"counseling-system/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler to handle the fake get request and db op
func Handler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Hello World, %s! You are %s, and %d years old", GetPerson().Name, GetPerson().Gender, GetPerson().Age)
	var res = Res{s, 200}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "JSON")
	result, _ := json.Marshal(res)

	fmt.Fprintf(w, string(result))

	db := utils.InitialDb()

	rows, err := db.Query("select username, password, phone from user")
	utils.CheckErr(err)

	for rows.Next() {
		var username string
		var password string
		var phone string

		err = rows.Scan(&username, &password, &phone)
		utils.CheckErr(err)
		fmt.Println(username, password, phone)
	}

}
