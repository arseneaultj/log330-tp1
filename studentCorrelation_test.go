
package main

import (
"testing"
)

func TestCT31(t *testing.T){
	students := []student{
		{"a",[6]float64{1,1,1,1,1,1},92},
		{"a",[6]float64{3,2,11,4,2,1},78},
		{"a",[6]float64{2,3,3,5,1,2},82},
		{"a",[6]float64{11,2,2,2,1.3,1.5},45}}

	correl := getStudentCorrelation(students,1)

	if correl < -1.999 && correl > -2.008 {
		t.Errorf("Expected -2.004 but got %v instead", correl)
	}
}

func TestCT32(t *testing.T){
	students := []student{
		{"a",[6]float64{1,50,1,1,1,1},98},
		{"a",[6]float64{3,1,11,4,2,1},12},
		{"a",[6]float64{2,40,3,5,1,2},97},
		{"a",[6]float64{11,4,2,2,1.3,1.5},25}}

	correl := getStudentCorrelation(students,2)
	if correl > -0.65 && correl < -0.66 {
		t.Errorf("Expected -0.65 but got %v instead", correl)
	}
}

//func TestCT33(t *testing.T){
//	students := []student{
//		{"a",[6]float64{1,50,1,-4,1,1},98},
//		{"a",[6]float64{3,1,11,4,-5,1},12},
//		{"a",[6]float64{2,40,3,5,-232.4,-2},97},
//		{"a",[6]float64{11,4,2,2,-1000,0.5},25}}
//
//	correl := getStudentCorrelation(students,-4)
//	if correl != math.NaN() {
//		t.Errorf("Expected NaN but got %v instead", correl)
//	}
//	}

