package main

import (
	"testing"
	"math"
)

func TestCT25(t *testing.T){
	testValue := line{1,5}
	yValue := 0.0
	x := testValue.getX(yValue)

	if x != -5 {
		t.Errorf("Expected -0.5 but got %v instead", x)
	}
}

func TestCT26(t *testing.T){
	testValue := line{-100, 999}
	yValue := -10000.0
	x := testValue.getX(yValue)

	if x != -899 {
		t.Errorf("Expected 109.99 but got %v instead", x)
	}
}

func TestCT27(t *testing.T){
	testValue := line{0, 0}
	yValue := 1.0
	x := testValue.getX(yValue)

	if !math.IsInf(x,1)  {
		t.Errorf("Expected NaN but got %v instead", x)
	}}


