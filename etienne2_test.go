
package main

import (
	"testing"
	"math"
)

func TestInvervalleHb(t *testing.T){

HighBoundValue := []point{
{546546543.0,123123123.0},
{246546543.0,323123123.0},
{346546543.0,423123123.0},
{146546543.0,523123123.0},
{346546543.0,523123123.0}}

intervalle := 	getInterval(HighBoundValue,1.86,1119)

	if math.Abs(intervalle- 3.2221271694698876e+08  ) > 0.001 /*Incertitude car comparaison de float*/  {
t.Errorf("Expected : %v Actual : %v ", 3.2221271694698876e+08,intervalle)
}
}



func TestInvervalleLb(t *testing.T){

HighBoundValue := []point{
{-546546543.0,-123123123.0},
{-246546543.0,-323123123.0},
{-346546543.0,-423123123.0},
{-146546543.0,-523123123.0},
{-346546543.0,-523123123.0}}

intervalle := 	getInterval(HighBoundValue,1.86,1119)

if math.Abs(intervalle- 3.222138264702489e+08  ) > 0.001 /*Incertitude car comparaison de float*/  {
t.Errorf("Expected : %v Actual : %v ", 3.222138264702489e+08,intervalle)
}
}





func TestInvervalleInvalid(t *testing.T){

invalidValues := []point{}

intervalle := 	getInterval(invalidValues,1.86,1119)

if intervalle == math.NaN() {
t.Errorf("No error hence test case failed")
}
}