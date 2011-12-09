package main

// Actions in gorl are represented as a floating point action value along with
// an integer id. Like states, the ids should be able to serve as indices into
// an array. In addition, we augment each action with a flag denoting whether
// the system currently belives the action is optimal for a given state.
type Action struct {
	Id     uint
	Val    float64
	Argmax bool
}
