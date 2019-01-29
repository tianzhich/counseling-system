package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// InitialDb initail the db connection and return the db instance
func InitialDb(dataSourceName string) *sql.DB {
	db, err := sql.Open("mysql", dataSourceName)
	CheckErr(err)

	return db
}
