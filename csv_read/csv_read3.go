package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/japanese"
)

func main() {
	path := "test.csv"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Open Error: %v", err)
	}
	defer file.Close()
	decoder := japanese.ShiftJIS.NewDecoder()
	reader := csv.NewReader(decoder.Reader(file))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			return
		}
		fmt.Println(record)
	}
}
