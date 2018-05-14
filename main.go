package main

import (
	"encoding/csv"
	"os"
	"bufio"
	"io"
	"fmt"
	"log"
)

func main() {

	readCsv(os.Args[1])
}

func readCsv(path string) [] int{

	file,err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}


	reader := csv.NewReader(bufio.NewReader(file))

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		fmt.Println(line)
	}

	return nil
}