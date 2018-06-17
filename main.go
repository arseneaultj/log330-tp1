package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"log"
	"strconv"
	"math"
)

type point struct{
	x float64
	y float64
}

func main() {

	values := readCsv(os.Args[1])

	moyenne := getAverage(values)
	variance := getVariance(values)
	ecartType := math.Sqrt(variance)
	fmt.Printf("Moyenne : %v\n", moyenne)
	fmt.Printf("Variance : %.4f\n", variance)
	fmt.Printf("Ecart type : %.2f\n", ecartType)

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

func getTotalSquaredDistance(values[]int) float64{

	var totalDist float64
	totalDist = 0
	average := getAverage(values)

	for _,value := range values {
		totalDist += (average-float64(value))*(average-float64(value))
	}

	return totalDist
}

func getVariance(values []int) float64{
	return getTotalSquaredDistance(values)*1/(float64(len(values)-1))
}