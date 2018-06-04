package main

import (
	"testing"
)

func TestCT7(t *testing.T){
	testValues := []int{1,1,1,1,1}
	variance := getVariance(testValues)

	if variance != 0 {
		t.Errorf("Expected 0 but got %v instead", variance)
	}
}

func TestCT8(t *testing.T){
	testValues := []int{-1000,-1000,0,1000,1000}
	variance := getVariance(testValues)

	if variance != 1000000 {
		t.Errorf("Expected 999999 but got %v instead", variance)
	}
}

func TestCT9(t *testing.T){

	testValues := []int{}
	variance := getVariance(testValues)

	if len(testValues) != 0 {
		t.Error("Expected array length of at least 1", variance)
	}
}

