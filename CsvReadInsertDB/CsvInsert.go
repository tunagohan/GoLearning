package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/encoding/japanese"
)

type Users struct {
	UserID   string
	Name     string
	Password string
	Sex      int
}

var i int = 0

func parsecsv(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	if i != 0 {
		s := reflect.ValueOf(v).Elem()
		recordlen := len(record)
		if s.NumField() != recordlen {
			log.Fatal("Field Miss Match: ", s.NumField(), " of ", recordlen)
		}
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			switch f.Type().String() {
			case "string":
				f.SetString(record[i])
			case "int":
				ival, err := strconv.ParseInt(record[i], 10, 0)
				if err != nil {
					log.Fatal("Parse Error: ", err)
				}
				f.SetInt(ival)
			default:
				log.Fatal("Unknown Type: ", f.Type)
			}
		}
	}
	return nil
}

func connectDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "{USER}:{PASS}@/{DB_NAME}")
	if err != nil {
		log.Fatal("Connecting Mysql Error: ", err)
	}
	return db
}

func main() {
	path := "{FILE_NAME}"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Open Error: ", err)
	}
	defer file.Close()

	decoder := japanese.ShiftJIS.NewDecoder()
	reader := csv.NewReader(decoder.Reader(file))
	reader.Comma = ','

	db := connectDB()
	insertsql, err := db.Prepare("INSERT INTO {TABLE_NAME} values(0,?,?,?,?)")
	if err != nil {
		log.Fatal("Prepare Error: ", err)
	}
	defer insertsql.Close()
	var users Users
	for {
		err := parsecsv(reader, &users)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if i != 0 {
			_, err = insertsql.Exec(users.UserID, users.Name, users.Password, users.Sex)
			if err != nil {
				log.Fatal("Exec SQL Error", err)
			}
			fmt.Printf("%s %s %s %d", users.UserID, users.Name, users.Password, users.Sex)
			fmt.Println(i, "---->insert")
		}
		i += 1
	}
	defer db.Close()
}
