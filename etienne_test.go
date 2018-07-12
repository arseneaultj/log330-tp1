package main

import (
	"testing"
	"math"
)

func TestCorrelationHb(t *testing.T){

highBoundValue := []point{
{546546543.0,123123123.0},
{246546543.0,323123123.0},
{346546543.0,423123123.0},
{146546543.0,523123123.0},
{346546543.0,523123123.0}}

correlation:= getCorrelation(highBoundValue);
// calculated at http://www.socscistatistics.com/tests/pearson/default2.aspx
if math.Abs(correlation- -0.7454) > 0.001 /*Incertitude car comparaison de float*/  {
t.Errorf("Expected : %v Actual : %v ", 0.5556,correlation)
}
}

func TestCorrelationLb(t *testing.T){

highBoundValue := []point{
{-546546543.0,-123123123.0},
{-246546543.0,-323123123.0},
{-346546543.0,-423123123.0},
{-146546543.0,-523123123.0},
{-346546543.0,-523123123.0}}

correlation:=getCorrelation(highBoundValue);
// calculated at http://www.socscistatistics.com/tests/pearson/default2.aspx
if math.Abs(correlation- -0.7454) > 0.001 /*Incertitude car comparaison de float*/  {
t.Errorf("Expected : %v Actual : %v ", 0.7454,correlation)
}
}




//Test a correlation where the denominator equals 0 (invalid division)

func TestCorrelationInvalid(t *testing.T){

invalidValues := []point{{0.0,0.0},{0.0,0.0},{0.0,0.0},{0.0,0.0}}

r := getCorrelation(invalidValues)

if math.Abs(r-1000.0) > 0.0001 /*Incertitude car comparaison de float (tolerence)*/  {
t.Errorf("Expected : %v Actual : %v ", 1000,r)
}
}
