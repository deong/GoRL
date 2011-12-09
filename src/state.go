package main

import (
	"math"
)

// States in gorl are represented by an array of coordinate values along with an
// integer id. The ids should be unique across the entire state space, and
// should serve as indices into an array of states (i.e., they should span the
// range [0, n-1].
type State struct {
	Id   uint
	Vals []float64
}

func MakeState(dim uint) (s State) {
	s = State{}
	s.Vals = make([]float64, dim)
	return
}

// compute the distance between two states
func EuclideanDistance(s1, s2 *State) (dist float64) {
	dist = 0.0
	for i := range s1.Vals {
		dist += (s1.Vals[i] - s2.Vals[i]) * (s1.Vals[i] - s2.Vals[i])
	}
	dist = math.Sqrt(dist)
	return
}