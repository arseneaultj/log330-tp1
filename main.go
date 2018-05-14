package main

import (
	"encoding/csv"
	"os"
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


	reader := csv.NewReader(file)

	values,_ := reader.ReadAll()

	for i := range values {
		fmt.Println(values[i])
	}

	//for {
	//	line,err := reader.Read()
	//
	//	if err == io.EOF {
	//		break
	//	} else if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	fmt.Println(line[0])
	//}
	return nil
}