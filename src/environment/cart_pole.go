package environment

import (
	"math"
	. "core"
)

type CartPoleEnv struct {
	steps uint
}

const (
	kForceMag       = 10.0
	kPoleLength     = 0.5 // action 1/2 the length
	kPoleMass       = 0.1
	kCartMass       = 1.0
	kTotalMass      = kPoleMass + kCartMass
	kPoleMassLength = kPoleMass * kPoleLength
	kGravity        = 9.8
	kTau            = 0.02 // seconds between updates
	kMaxSteps       = 1000
)

var cpFeatureRanges = []Range{
	Range{-4.5, 4.5},
	Range{-3.0, 3.0},
	Range{-math.Pi / 4.0, math.Pi / 4.0},
	Range{-math.Pi / 4.0, math.Pi / 4.0},
}

func (env *CartPoleEnv) Features() (f []Range) {
	f = cpFeatureRanges
	return
}

func (env *CartPoleEnv) ActionRange() Range {
	return Range{-1.0, 1.0}
}

// return the immediate reward and new state resulting from taking action a in state s
func (env *CartPoleEnv) ApplyAction(s State, a Action) (newState State, reward float64) {
	force := a.Val * kForceMag
	cosTheta := math.Cos(s.Vals[2])
	sinTheta := math.Sin(s.Vals[2])
	temp := (force + kPoleMassLength*s.Vals[3]*s.Vals[3]*sinTheta) / kTotalMass
	deltaTheta := (kGravity*sinTheta - cosTheta*temp) /
		(kPoleLength * (4.0/3.0 - kPoleMass*cosTheta*cosTheta/kTotalMass))
	deltaX := temp - kPoleMassLength*deltaTheta*cosTheta/kTotalMass

	// compute the next state
	newState = MakeState(4)
	newState.Vals[0] = s.Vals[0] + kTau*s.Vals[1]
	newState.Vals[1] = s.Vals[1] + kTau*deltaX
	newState.Vals[2] = s.Vals[2] + kTau*s.Vals[3]
	newState.Vals[3] = s.Vals[3] + kTau*deltaTheta

	// calculate the reward
	reward = 0
	if env.AtFailState(newState) {
		reward = -1
	}

	env.steps++
	return
}

// check if we're at a goal state
func (env *CartPoleEnv) AtGoalState(s State) bool {
	if env.steps >= kMaxSteps {
		return true
	}
	return false
}

// check if we're at the fail state
func (env *CartPoleEnv) AtFailState(s State) bool {
	if math.Fabs(s.Vals[0]) > 4.0 || math.Fabs(s.Vals[2]) > math.Pi/4.0 {
		return true
	}
	return false
}

// return the start state
func (env *CartPoleEnv) StartState() (s State) {
	s = State{0, []float64{0.0, 0.0, 0.0, 0.0}}
	return
}

// reset the count
func (env *CartPoleEnv) Reset() {
	env.steps = 0
}