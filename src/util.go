package main

import (
	"container/list"
)

// return a slice of evenly spaced points
func Linspace(minVal, maxVal float64, nPoints int) []float64 {
	v := make([]float64, nPoints)
	curVal, delta := minVal, (maxVal - minVal) / float64(nPoints - 1)
	for i := 0; i < nPoints; i++ {
		v[i] = curVal
		curVal += delta
	}
	return v
}


// Build a lattice from a given set of arbitrarily spaced dimensions.
// 
// For example, given dimensions [a, b, c] and [x, y], the lattice
// constructed should look like
//
//     [[a, x],
//	    [a, y],
//      [b, x],
//      [b, y],
//      [c, x],
//      [c, y]]
//
// Each dimension adds one coordinate to the individual rows in the output
// vector, and the length of each dimension controls the number of different
// values along that dimension that will be generated in the output.
//
// The function works as follows: suppose you have three lattice points [-1, 0,
// 1], and you wish to build the lattice in three dimensions. The desired
// output should be a list of 27 points,
// 
//     [[-1, -1, -1],
//      [-1, -1,  0],
//      [-1, -1,  1],
//      [-1,  0, -1],
//      ...
//      [ 1,  1,  1]]
//
// This function builds that list via the following idea. First, you create a
// vector of three partial states, like such [-1], [0], and [1], and then we
// push those onto a stack. Then, we pop one off, and if it doesn't contain the
// required number of dimensions already, we create three new points by
// appending every possible lattice point onto the popped point. So if we pop
// [1], we get [1, -1], [1, 0], and [1, 1].  Each is pushed onto the partially
// completed stack. When a point is popped off that is complete, it is added to
// the list of states. When the stack is empty, all the states have been
// created.
func BuildLattice(dimensions [][]float64) []State { 
	// calculate the total number of states and init some memory for them
	n := 1 
	for i := range dimensions { 
		n *= len(dimensions[i]) 
	} 
	states := make([]State, n)
	partial := list.New()

	// create some temp storage for partially constructed states
	for i := 0; i < len(dimensions[0]); i++ {
		s := new(State)
		s.vals = make([]float64, 0)
		s.vals = append(s.vals, dimensions[0][i])
		partial.PushBack(*s)
	}

	index := 0
	for partial.Len() > 0 {
		p := partial.Front()
		s := p.Value.(State)
		partial.Remove(p)
		
		// if the newly grabbed state is finally constructed, push
		// it onto the states array. otherwise, make enough copies
		// to append all possible values of the next dimension and
		// put them back onto the partial array.
		if len(s.vals) == len(dimensions) {
			s.id = index
			states[index] = s
			index++
		} else {
			for i := 0; i < len(dimensions[len(s.vals)]); i++ {
				sp := new(State)
				sp.vals = make([]float64, len(s.vals))
				copy(sp.vals, s.vals)
				sp.vals = append(sp.vals, dimensions[len(sp.vals)][i])
				partial.PushBack(*sp)
			}
		}
	}
	return states
}

