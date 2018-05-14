package main

import (
	"testing"
	"math"
)

func TestMoyenne(t *testing.T){
	testValues := []int{10,20,30,40,50}
	average := getAverage(testValues)

	if average != 30 {
		t.Errorf("Expected 30 but got %v instead", average)
	}
}

func TestSommeDistanceCarre(t *testing.T){
	testValues := []int{10,20,30,40,50}
	total := getTotalSquaredDistance(testValues)

	if total != 1000 {
		t.Errorf("Expected 60 but got %v instead", total)
	}
}

func TestVariance(t *testing.T){
	testValues := []int{10,20,30,40,50}
	variance := getVariance(testValues)

	if variance != 250 {
		t.Errorf("Expected 250 but got %v instead", variance)
	}
}

func TestEcartType (t *testing.T){
	if math.Sqrt(25) != 5 {
		t.Error("Expected 5 but got something else")
	}
}