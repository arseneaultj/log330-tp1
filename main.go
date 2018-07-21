package main

import (
	"encoding/csv"
	"os"
	"fmt"
	"log"
	"strconv"
	"math"
	"bufio"
	"strings"
)

const alpha70, alpha90 = 1.108, 1.860
const xk = 1119

type point struct{
	x float64
	y float64
}

type line struct{
	slope float64
	constant float64
}

type student struct{
	name string
	effort [6]float64
	grade float64
}

func main() {
	values := readCsv(os.Args[1])
	line := line{getSlope(values),getConstant(values)}
	interval70 := getInterval(values,alpha70,1119)
	interval90 := getInterval(values,alpha90,1119)
	predictionLoc := line.getY(xk)
	fmt.Printf("Interval(α=0.9) : [%.5f - %.5f] \nInterval(α=0.7) : [%.5f - %.5f] \n ",
			predictionLoc-interval90, predictionLoc+interval90, predictionLoc-interval70, predictionLoc+interval70)

}

func getStdDev(points []point) float64{
	line := line{getSlope(points),getConstant(points)}
	sum := 0.0
	for _,point := range points {
		dist := point.y - line.getY(point.x)
		sum += dist*dist
	}
	return math.Sqrt(1/float64(len(points)-1)*sum)
}

func getInterval(points []point, alpha float64,xk int)float64{
	avgX := getFAverage(getAllX(points))
	top := float64(xk)-avgX

	interval := alpha*getStdDev(points)*math.Sqrt(1 + 1.0/float64(len(points)) + (top*top)/getTotalFSquaredDistance(getAllX(points)))
	return interval
}

func readUserInput() [2]float64{
	reader := bufio.NewReader(os.Stdin)
	choice := ""
	for strings.Compare(choice, "x") != 0 && strings.Compare(choice, "y") != 0{
		fmt.Print("Choose variable to input(x or y) : ")
		val,_ := reader.ReadString('\n')
		choice = strings.Replace(val, "\r\n", "", -1)
	}

	value,err := strconv.ParseFloat("a",64) //generate error
	for err != nil {
		fmt.Println("Enter value for " + choice)
		val,_ := reader.ReadString('\n')
		val = strings.Replace(val, "\r\n", "", -1)
		value,err = strconv.ParseFloat(val,64)
	}

	if choice == "x" {
		return [2]float64{0,value}
	} else {
		return [2]float64{1,value}
	}
}

func getAllX(points []point)[]float64{
	var values []float64
	for _,point := range points {
		values = append(values,point.x)
	}
	return values
}

func (line line) getY(x float64) float64{
	return line.slope*x + line.constant
}

func (line line)getX(y float64) float64{
	return y/line.slope - line.constant
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

func getConstant(points []point) float64{
	averages := getPointsAverage(points)
	slope := getSlope(points)

	return averages[1] - slope*averages[0]
}

func getStudentCorrelation(student []student, week int) float64{
	var points []point
	for _,stud := range student {
		points = append(points, studentToPoint(stud,week))
	}

	return  getCorrelation(points)
}

func studentToPoint(student student, week int) point{
	return point{student.effort[week],student.grade}
}

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


func readCsv(path string) []point{

	var values []point

	file,err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	lines,_ := reader.ReadAll()

	for i := range lines {
		x,_:= strconv.ParseFloat(lines[i][0],64)
		y,_:= strconv.ParseFloat(lines[i][1],64)
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