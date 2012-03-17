package main

import (
	"testing"
	"math"
)

type linspaceTest struct {
	minVal  float64
	maxVal  float64
	nPoints int
	result  []float64
}

var linspaceTests = []linspaceTest{
	linspaceTest{-1.0,   1.0,  5, []float64{-1.0, -0.5, 0.0, 0.5, 1.0}},
	linspaceTest{ 0.0,  10.0,  4, []float64{0.0, 3.333333, 6.666667, 10.0}},
	linspaceTest{ 2.0,   3.0,  2, []float64{2.0, 3.0}},
	linspaceTest{ 0  ,   1  , 11, []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0}},
	linspaceTest{-2.0,  -1.0,  3, []float64{-2.0, -1.5, -1.0}},
}

func TestLinspace(t *testing.T) {
	for _, lt := range linspaceTests {
		val := Linspace(lt.minVal, lt.maxVal, lt.nPoints)
		if !vectorEpsilonEqual(val, lt.result, 0.00001) {
			t.Errorf("Error: %v != %v\n", val, lt.result)
		}
	}
}

func vectorEpsilonEqual(v1, v2 []float64, epsilon float64) bool {
	for i := range v1 {
		if !epsilonEqual(v1[i], v2[i], epsilon) {
			return false
		}
	}
	return true
}

func epsilonEqual(x, y, epsilon float64) bool {
	if math.Abs(x-y) < epsilon {
		return true
	}
	return false
}
