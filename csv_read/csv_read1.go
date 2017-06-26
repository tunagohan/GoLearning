//SJIS用
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	path := "test.csv"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("open Error: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	reader.Comma = '\t'
	reader.LazyQuotes = true

	for {

		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("read Error: %v", err)
		}
		fmt.Println(record)
	}
}
