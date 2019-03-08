package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

const dataSourceName = "tianzhi:tianzhi@tcp(47.94.223.143:3306)/pcs"

func initialDb() *sql.DB {
	mydbcon := dataSourceName
	db, err := sql.Open("mysql", mydbcon)
	CheckErr(err)
	return db
}

// QueryDB to Query the db
func QueryDB(str string) *sql.Rows {
	var db = initialDb()
	defer db.Close()

	rows, err := db.Query(str)
	CheckErr(err)

	return rows
}

// QueryDBRow to query the number of row
func QueryDBRow(str string) int {
	var count int
	var db = initialDb()
	defer db.Close()

	err := db.QueryRow(str).Scan(&count)
	CheckErr(err)

	return count
}

// InsertDB to insert data to db
func InsertDB(str string, args ...interface{}) (int64, bool) {
	var db = initialDb()
	defer db.Close()

	stmt, err := db.Prepare(str)
	CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	CheckErr(err)

	if rows, _ := res.RowsAffected(); rows == 1 {
		rowID, _ := res.LastInsertId()
		return rowID, true
	}

	return -1, false
}

// UpdateDB to update data in db
func UpdateDB(str string, args ...interface{}) bool {
	var db = initialDb()
	defer db.Close()

	stmt, err := db.Prepare(str)
	CheckErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	CheckErr(err)

	if rows, _ := res.RowsAffected(); rows == 1 {
		return true
	}

	return false
}
