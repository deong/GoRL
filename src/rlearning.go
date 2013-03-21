package main

import (
	"math/rand"
	"os"
	"fmt"
	"math"
)

type RLearning struct {
	states  []State
	actions []Action
	Q       [][]float64
	rho     float64
	alpha   float64
	beta    float64
	epsilon float64
}

// Initialize the Q-values table and trace.
func (self *RLearning) Init(env Environment) {
	// create the state space
	var nPoints []int
	var err error
	if nPoints, err = IntArrayParameter("environment", "state_grid"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	grid := make([][]float64, len(nPoints))
	featureRanges := env.Features()
	for fi := range featureRanges {
		grid[fi] = Linspace(featureRanges[fi].Min, featureRanges[fi].Max, nPoints[fi])
	}
	self.states = BuildLattice(grid)
	numStates := len(self.states)

	// create the action space
	actionRange := env.ActionRange()
	var aPoints uint
	if aPoints, err = UintParameter("environment", "action_grid"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	aGrid := Linspace(actionRange.Min, actionRange.Max, int(aPoints))
	self.actions = make([]Action, len(aGrid))
	for ai := range aGrid {
		self.actions[ai].Id = uint(ai)
		self.actions[ai].Val = aGrid[ai]
	}
	numActions := len(self.actions)

	// initialize the Q-values array and execution trace
	self.Q = make([][]float64, numStates)
	for i := 0; i < numStates; i++ {
		self.Q[i] = make([]float64, numActions)
	}

	// initialize the average reward
	self.rho = 0

	// set up some learning parameters
	if self.alpha, err = Float64Parameter("learning", "alpha"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if self.beta, err = Float64Parameter("learning", "beta"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if self.epsilon, err = Float64Parameter("learning", "epsilon"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Return the index of the best action from a given state
func (self *RLearning) ArgmaxAction(s State) (indexOfBest uint, valueOfBest float64) {
	indexOfBest = 0
	valueOfBest = self.Q[s.Id][0]
	for i := 1; i < len(self.Q[s.Id]); i++ {
		if self.Q[s.Id][i] > valueOfBest {
			indexOfBest, valueOfBest = uint(i), self.Q[s.Id][i]
		}
	}
	return
}

// Return a random action and its estimated value
func (self *RLearning) RandomAction(s State) (indexOfBest uint, valueOfBest float64) {
	indexOfBest = uint(rand.Intn(len(self.Q[s.Id])))
	valueOfBest = self.Q[s.Id][indexOfBest]
	return
}

// Return an epsilon-greedy action, its estimated value, and whether it was chosen greedily
func (self *RLearning) EpsilonGreedyAction(s State, epsilon float64) (indexOfBest uint, valueOfBest float64, wasGreedy bool) {
	if rand.Float64() < epsilon {
		indexOfBest, valueOfBest = self.RandomAction(s)
		wasGreedy = false
	} else {
		indexOfBest, valueOfBest = self.ArgmaxAction(s)
		wasGreedy = true
	}
	return
}

// Learn the Q-values
func (self *RLearning) Learn(env Environment) {
	s := env.StartState()
	self.DiscretizeState(&s)
	numSteps := 0
	for {
		// select an action
		aIndex, _, _ := self.EpsilonGreedyAction(s, self.epsilon)
		a := self.actions[aIndex]

		// observe reward, next state
		sp, reward := env.ApplyAction(s, a)
		self.DiscretizeState(&sp)

		// calculate optimum value of next action
		_, argmaxAP := self.ArgmaxAction(sp)

		// calculate error
		delta := reward - self.rho + argmaxAP - self.Q[s.Id][a.Id]

		// update the policy
		self.Q[s.Id][a.Id] += self.alpha * delta
		if math.Abs(self.Q[s.Id][a.Id] - argmaxAP) < 1e-8 {
			self.rho += self.beta * delta
		}

		// iterate the policy
		s = sp
		numSteps++

		// reduce the exploration rate
		if numSteps%100 == 0 {
			self.epsilon *= 0.95
		}
	}
}

// given an arbitrary state vector, set its id to that of the nearest state in the space
func (self *RLearning) DiscretizeState(s *State) {
	// TODO: do a more efficient calculation to replace this search
	idOfNearest := 0
	distToNearest := math.MaxFloat64
	for i := range self.states {
		currentDist := EuclideanDistance(s, &self.states[i])
		if currentDist < distToNearest {
			idOfNearest = i
			distToNearest = currentDist
		}
	}
	s.Id = uint(idOfNearest)
}


func (self *RLearning) FollowPolicy(env Environment) {
	env.Reset()
	s := env.StartState()
	num_steps := 0
	for !env.AtGoalState(s) && !env.AtFailState(s) {
		self.DiscretizeState(&s)
		// select an action
		aIndex, _ := self.ArgmaxAction(s)
		a := self.actions[aIndex]
		fmt.Printf("step %d: executing action %v\n", num_steps+1, a.Val)
		num_steps++
		// observe reward, next state
		s, _ = env.ApplyAction(s, a)
	}
}
