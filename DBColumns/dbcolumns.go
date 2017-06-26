package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getDBError(err error) {
	if err != nil {
		log.Fatal("DB Error: %v", err)
	}
}

func columnsError(err error) {
	if err != nil {
		log.Fatal("columns Error: %v", err)
	}
}

func getDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "{USER}:{PASS}@/{DB_NAME}")
	getDBError(err)
	return db
}

func columns(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM {TABLE_NAME}")
	columnsError(err)
	columns, err := rows.Columns()
	columnsError(err)
	values := make([]sql.RawBytes, len(columns))
	for i := range values {
		fmt.Println(columns[i])
	}
}

func main() {
	fmt.Println("** mysql start ** ")
	db := getDB()
	columns(db)
	defer db.Close()
	fmt.Println("** mysql end ** ")
}
