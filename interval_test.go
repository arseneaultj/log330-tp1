package main

import (
	"testing"
	"math"
)

func TestCT37(t *testing.T){
	testValue := []point{{1.3434,3.423543},{2.4332,4.111},{2.2222,6.666},{1.111,9.9999},{3.333,3.3333}}
	interval := getInterval(testValue,alpha70,1000)

	if toFixed(interval, 4) != 1442.5143 {
		t.Errorf("Expected 1442.5143 but got %v instead", interval)
	}
}

func TestCT38(t *testing.T){
	testValue := []point{{99.9999,111.11111},{22222.4332,3333.111},{44444.2222,5555.666},
		{55555.111,666.9999},{5.333,1.3333}}
	interval := getInterval(testValue,alpha90,1500)

	if toFixed(interval, 4) != 4675.0347 {
		t.Errorf("Expected 4675.0347 but got %v instead", interval)
	}
}

func TestCT39(t *testing.T){
	testValue := []point{{1,1},{1,1}}
	interval := getInterval(testValue,alpha90,1)

	if !math.IsNaN(interval) {
		t.Errorf("Expected NaN but got %v instead", interval)
	}
}
