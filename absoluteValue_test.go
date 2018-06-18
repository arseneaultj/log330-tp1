package main

import (
	"testing"
)

func TestCT10(t *testing.T){
	testValue := -9999.2
	absValue := absolute(testValue)

	if absValue != 9999.2 {
		t.Errorf("Expected 9999.2 but got %v instead", absValue)
	}
}

func TestCT11(t *testing.T){
	testValue := 3828566.4
	absValue := absolute(testValue)

	if absValue != 3828566.4 {
		t.Errorf("Expected 3828566.4 but got %v instead", absValue)
	}
}

func TestCT12(t *testing.T){

	testValue := 0.0
	absValue := absolute(testValue)

	if absValue != 0 {
		t.Errorf("Expected 0 but got %v instead", absValue)
	}
}