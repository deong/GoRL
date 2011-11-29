package environment

import (
	"math"
	. "core"
)

type MountainCarEnv struct {
	steps uint
}

var mcFeatureRanges = []Range{
	Range{-1.4, 1.4}, 			// cart position
	Range{-0.7, 0.7},			// cart velocity
}

func (env *MountainCarEnv) Features() (f []Range) {
	f = mcFeatureRanges
	return
}

func (env *MountainCarEnv) ActionRange() Range {
	return Range{-1.0, 1.0}
}

// return the immediate reward and new state resulting from taking action a in state s
func (env *MountainCarEnv) ApplyAction(s State, a Action) (newState State, reward float64) {
	// compute the new velocity
	cartVelocity := s.Vals[1] + (0.001 * a.Val) + (-0.0025 * math.Cos(3.0 * s.Vals[0]))
	
	// truncate at boundaries
	if cartVelocity < mcFeatureRanges[1].Min {
		cartVelocity = mcFeatureRanges[1].Min
	} else if cartVelocity > mcFeatureRanges[1].Max {
		cartVelocity = mcFeatureRanges[1].Max
	}

	// compute the new position
	cartPosition := s.Vals[0] + cartVelocity
	
	// truncate at boundaries
	if cartPosition < mcFeatureRanges[0].Min {
		cartPosition = mcFeatureRanges[0].Min
	} else if cartPosition > mcFeatureRanges[0].Max {
		cartPosition = mcFeatureRanges[0].Max
	}

	newState = MakeState(2)
	newState.Vals[0] = cartPosition
	newState.Vals[1] = cartVelocity

	reward = -1.0
	if env.AtGoalState(newState) {
		reward = 100.0
	}

	return
}

// check if we're at a goal state
func (_ *MountainCarEnv) AtGoalState(s State) bool {
	if s.Vals[0] >= 0.45 {
		return true
	}
	return false
}

// There is no fail state for the mountain car problem
func (_ *MountainCarEnv) AtFailState(_ State) bool {
	return false
}

// return a start state
func (_ *MountainCarEnv) StartState() (s State) {
	s = State{0, []float64{-0.5, 0.0}}
	return
}

// reset the environment (nothing to do for this problem)
func (_ *MountainCarEnv) Reset() {
	return
}