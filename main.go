package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"log"
	"strconv"
)

func main() {

	data := readCsv(os.Args[1])
	for d := range data {
		fmt.Println(data[d])
	}

	average := getAverage(data)
	fmt.Println(average)

}

func readCsv(path string) [] int{

	var values []int

	file,err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}



	reader := csv.NewReader(file)

	lines,_ := reader.ReadAll()

	for i := range lines {
		value,_ := strconv.Atoi(lines[i][0])
		values = append(values, value)
	}

	return values
}

func getAverage(values []int) float64{

	total := 0
	for _,value := range values{
		total += value
	}
	return float64(total)/float64(len(values))
}