package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// CheckErr to check err like db, file operation
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// InitialDb initail the db connection and return the db instance
func InitialDb() *sql.DB {
	mydbcon := dataSourceName
	db, err := sql.Open("mysql", mydbcon)
	CheckErr(err)
	return db
}
