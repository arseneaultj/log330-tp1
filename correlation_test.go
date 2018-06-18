package main

import (
	"testing"
	"math"
)

func TestCT16(t *testing.T){
	testValue := []point{{10,1},{1,10}}
	correlation := getCorrelation(testValue)

	if absolute(correlation) != 1 {
		t.Errorf("Expected 1 but got %v instead", correlation)
	}
}

func TestCT17(t *testing.T){
	testValue := []point{{164560,54646},{65346,23352352},{32423432, 52542}}
	correlation := getCorrelation(testValue)

	if absolute(correlation) < 0.5 || absolute(correlation) > 0.6 {
		t.Errorf("Expected 0.5023 but got %v instead", correlation)
	}
}

func TestCT18(t *testing.T){
	testValue := []point{{1,1},{1,1}}
	correlation := getCorrelation(testValue)

	if !math.IsNaN(correlation)  {
		t.Errorf("Expected NaN but got %v instead", correlation)
	}}
