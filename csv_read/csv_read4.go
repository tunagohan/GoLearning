package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

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
	reader := bufio.NewReader(decoder.Reader(file))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Read Error: %v", err)
		}
		cols := strings.Split(string(line), ",")
		fmt.Println(cols)
	}
}
