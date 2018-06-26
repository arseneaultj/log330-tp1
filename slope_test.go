package main

import (
	"testing"
	"math"
)

func TestCT19(t *testing.T){
	testValues := []point{{0,0},{10,10},{5,5},{100,100}}
	slope := getSlope(testValues)

	if slope != 1 {
		t.Errorf("Expected 1 but got %v instead", slope)
	}
}

func TestCT20(t *testing.T){
	testValues := []point{{0,0},{-10,100},{10, -100}}
	slope := getSlope(testValues)

	if slope != -10 {
		t.Errorf("Expected -10 but got %v instead", slope)
	}
}

func TestCT21(t *testing.T){
	testValues := []point{{1,1}}
	slope := getSlope(testValues)

	if !math.IsNaN(slope)  {
		t.Errorf("Expected NaN but got %v instead", slope)
	}}
