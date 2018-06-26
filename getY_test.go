package main

import (
	"testing"
	"math"
)

func TestCT28(t *testing.T){
	testValue := line{1,5}
	xValue := 0.0
	y := getY(xValue,testValue)

	if y != 5 {
		t.Errorf("Expected 5 but got %v instead", y)
	}
}

func TestCT29(t *testing.T){
	testValue := line{-100, 999}
	xValue := -10000.0
	y := getY(xValue,testValue)

	if y != 1000999 {
		t.Errorf("Expected 1000999 but got %v instead", y)
	}
}

func TestCT30(t *testing.T){
	testValue := line{0, 0}
	xValue := math.Inf(1)
	y := getY(xValue,testValue)

	if !math.IsNaN(y)  {
		t.Errorf("Expected NaN but got %v instead", y)
	}}
