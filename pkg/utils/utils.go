package utils

import "net/http"

// CheckErr to check err like db, file operation
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// AllowCors to config the CORS
func AllowCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Add("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Content-Type", "application/json")
}
