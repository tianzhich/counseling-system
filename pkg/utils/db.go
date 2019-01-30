package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

const dsn = "tianzhi:tianzhi@tcp(47.94.223.143:3306)/pcs"

// InitialDb initail the db connection and return the db instance
func InitialDb() *sql.DB {
	db, err := sql.Open("mysql", dsn)
	CheckErr(err)

	return db
}
