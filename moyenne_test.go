package main

import (
	"testing"
)

func TestCT1(t *testing.T){
	testValues := []int{-999999,-999999,-999999,-999999,-999999}
	average := getAverage(testValues)

	if average != -999999 {
		t.Errorf("Expected -999999 but got %v instead", average)
	}
}

func TestCT2(t *testing.T){
	testValues := []int{999999,999999,999999,999999,999999}
	average := getAverage(testValues)

	if average != 999999 {
		t.Errorf("Expected 999999 but got %v instead", average)
	}
}

func TestCT3(t *testing.T){

	testValues := []int{}
	average := getAverage(testValues)

	if len(testValues) != 0 {
		t.Error("Expected array length of at least 1", average)
	}
}
