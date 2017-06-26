// SJISをUTF-8に変換
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
		log.Fatal("path error: %v", err)
	}
	defer file.Close()

	decoder := japanese.ShiftJIS.NewDecoder()
	reader := csv.NewReader(decoder.Reader(file))
	// ダブルクォートを調べない
	reader.Comma = '\t'
	reader.LazyQuotes = true
	log.Printf("Start")
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("read error: %v", err)
		}
		fmt.Println(record)
	}
	log.Printf("End")
}
