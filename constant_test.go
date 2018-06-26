package main

import (
	"testing"
	"math"
)

func TestCT22(t *testing.T){
	testValues := []point{{0,0},{10,10},{5,5},{100,100}}
	cst := getConstant(testValues)

	if cst != 0 {
		t.Errorf("Expected 0 but got %v instead", cst)
	}
}

func TestCT23(t *testing.T){
	testValues := []point{{0,1000},{-10,1100},{10, 900}}
	cst := getConstant(testValues)

	if cst != 1000 {
		t.Errorf("Expected 1000 but got %v instead", cst)
	}
}

func TestCT24(t *testing.T){
	testValues := []point{{1,1}}
	cst := getConstant(testValues)

	if !math.IsNaN(cst)  {
		t.Errorf("Expected NaN but got %v instead", cst)
	}}

