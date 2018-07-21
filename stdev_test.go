package main

import (
"testing"
"math"
)

func TestCT34(t *testing.T){
testValue := []point{{1.3434,3.423543},{2.4332,4.111},{2.2222,6.666},{1.111,9.9999},{3.333,3.3333}}
stdev := getStdDev(testValue)

if toFixed(stdev, 4) != 2.3324 {
t.Errorf("Expected 2.3324 but got %v instead", stdev)
}
}

func TestCT35(t *testing.T){
testValue := []point{{99.9999,111.11111},{22222.4332,3333.111},{44444.2222,5555.666},
{55555.111,666.9999},{5.333,1.3333}}
stdev := getStdDev(testValue)

if toFixed(stdev, 4) != 2119.8737 {
	t.Errorf("Expected 2119.8737 but got %v instead", stdev)
		}
}

func TestCT36(t *testing.T){
testValue := []point{{1,1},{1,1}}
stdev := getStdDev(testValue)

if !math.IsNaN(stdev) {
	t.Errorf("Expected NaN but got %v instead", stdev)
	}
}
