package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"

	"golang.org/x/text/encoding/japanese"
)

type Test struct {
	ID       int
	User_ID  string
	Name     string
	Password string
	Sex      int
}

func parsecsv(reader *csv.Reader, v interface{}) error {
	record, err := reader.Read()
	if err != nil {
		return err
	}
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return &fieldmissmatch{s.NumField(), len(record)}
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return &unknowtype{f.Type().String()}
		}
	}
	return nil
}

func main() {
	path := "test.csv"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Open Error: %v", err)
	}
	defer file.Close()
	decoder := japanese.ShiftJIS.NewDecoder()
	reader := csv.NewReader(decoder.Reader(file))

	reader.Comma = ','
	var test Test
	for {
		err := parsecsv(reader, &test)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%d %s %s %s %d\n", test.ID, test.User_ID, test.Name, test.Password, test.Sex)
	}
}

type fieldmissmatch struct {
	expected, found int
}

func (e *fieldmissmatch) Error() string {
	return "CSV fields mismatch: " + strconv.Itoa(e.expected) + " of " + strconv.Itoa(e.found)
}

type unknowtype struct {
	Type string
}

func (e *unknowtype) Error() string {
	return "Unsupported type: " + e.Type
}
