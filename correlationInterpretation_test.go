package main

import (
	"testing"
)

func TestCT13(t *testing.T){
	testValue := -1.0
	interpretation := evaluateCorrelation(testValue)

	if interpretation != "Très forte à parfaite" {
		t.Errorf("Expected Très forte à parfaite but got %v instead", interpretation)
	}
}

func TestCT14(t *testing.T){
	testValue := 0.2
	interpretation := evaluateCorrelation(testValue)

	if interpretation != "Faible à moyenne" {
		t.Errorf("Expected Faible à moyenne but got %v instead", interpretation)
	}
}

func TestCT15(t *testing.T){

	testValue := 1.5
	interpretation := evaluateCorrelation(testValue)

	if interpretation != "Correlation invalide" {
		t.Errorf("Expected Correlation invalide but got %v instead", interpretation)
	}
}
