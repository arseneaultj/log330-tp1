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
	fmt.Println(getSlope(values))
	//correlation := getCorrelation(values)
	//fmt.Printf("Correlation : %.8f\n",correlation)
	//fmt.Println("La correlation est " + evaluateCorrelation(correlation))
}
func getPointsAverage(points []point) [2]float64{
	var xValues []float64
	var yValues []float64

	for _,val := range points {
		xValues = append(xValues, val.x)
		yValues = append(yValues, val.y)
	}
	return [2]float64{getFAverage(xValues),getFAverage(yValues)}
}

func getSlope(points []point) float64 {
	averages := getPointsAverage(points)
	n := float64(len(points))

	xySum := 0.0
	x2sum := 0.0

	for _,val := range points {
		xySum += val.x*val.y
		x2sum += val.x * val.x
	}

	top := xySum - n*averages[0]*averages[1]
	bottom := x2sum - n*averages[0]*averages[0]

	return top/bottom
}

//func getConstant(points []point) float64{
//
//}

func getCorrelation(points []point) float64{
	n := float64(len(points))
	var xValues []float64
	var yValues []float64
	xSum := 0.0
	ySum := 0.0
	x2Sum := 0.0
	y2Sum := 0.0

	top := 0.0

	for _,value := range points {
		xValues = append(xValues, value.x)
		yValues = append(yValues, value.y)
		xSum += value.x
		ySum += value.y
		x2Sum += value.x*value.x
		y2Sum += value.y*value.y

		}
	xAvg := getFAverage(xValues)
	yAvg := getFAverage(yValues)

	for _,value := range points {
		top += (value.x-xAvg)*(value.y-yAvg)
	}

	bottom := (n*x2Sum-(xSum*xSum))*(n*y2Sum-(ySum*ySum))


	r := (n*top)/math.Sqrt(bottom)

	return r

}

func absolute(value float64) float64{
	if value < 0 {
		value *= -1
	}
	return value
}

func evaluateCorrelation(correl float64) string{

	r := absolute(correl)

	if r < 0.2 {
		return "Nulle ou faible"
	}
	if r < 0.4 {
		return "Faible à moyenne"
	}
	if r < 0.7 {
		return "Moyenne à forte"
	}
	if r < 0.9 {
		return "Forte à très forte"
	}
	if r <= 1{
		return "Très forte à parfaite"
	}else {
		return "Correlation invalide"
	}
}

func readCsv(path string) [] point{

	var values []point

	file,err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	lines,_ := reader.ReadAll()

	for i := range lines {
		x,_ := strconv.ParseFloat(lines[i][0],64)
		y,_ := strconv.ParseFloat(lines[i][1],64)
		values = append(values, point{x,y})
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

func getFAverage(values []float64) float64{

	total := 0.0
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

func getTotalFSquaredDistance(values[]float64) float64{

	var totalDist float64
	totalDist = 0
	average := getFAverage(values)

	for _,value := range values {
		totalDist += (average-value)*(average-value)
	}

	return totalDist
}

func getTotalDistance(values[]float64) float64{

	var totalDist float64
	totalDist = 0
	average := getFAverage(values)

	for _,value := range values {
		totalDist += absolute(average-value)
	}

	return totalDist
}

func getVariance(values []int) float64{
	return getTotalSquaredDistance(values)*1/(float64(len(values)-1))
}