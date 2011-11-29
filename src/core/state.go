package core

// States in gorl are represented by an array of coordinate values along with an
// integer id. The ids should be unique across the entire state space, and
// should serve as indices into an array of states (i.e., they should span the
// range [0, n-1].
type State struct {
	Id   uint
	Vals []float64
}


