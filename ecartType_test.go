package main

import (
"testing"
	"math"
)

func TestCT4(t *testing.T){
	if math.Sqrt(1) != 1 {
		t.Error("Expected 1 but got something else")
	}
}

func TestCT5(t *testing.T){
	if math.Sqrt(10000000000) != 100000 {
		t.Error("Expected 100000 but got something else")
	}
}

func TestCT6(t *testing.T){

	testValues := []int{}
	variance := getVariance(testValues)

	if len(testValues) != 0 {
		t.Error("Expected array length of at least 1", variance)
	}
}
