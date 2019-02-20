package utils

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

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

// InitialDb initail the db connection and return the db instance
func InitialDb() *sql.DB {
	mydbcon := dataSourceName
	db, err := sql.Open("mysql", mydbcon)
	CheckErr(err)
	return db
}
