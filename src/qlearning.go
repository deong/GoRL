package main

import (
	"math/rand"
	"os"
	"fmt"
	"math"
)

type QLearning struct {
	states    []State
	actions   []Action
	Q         [][]float64
	E         [][]float64
	maxEpochs uint
	alpha     float64
	gamma     float64
	lambda    float64
	epsilon   float64
}

// Initialize the Q-values table and trace.
func (self *QLearning) Init(env Environment) {
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
	self.E = make([][]float64, numStates)
	for i := 0; i < numStates; i++ {
		self.Q[i] = make([]float64, numActions)
		self.E[i] = make([]float64, numActions)
	}

	// set up some learning parameters
	if self.maxEpochs, err = UintParameter("learning", "epochs"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if self.alpha, err = Float64Parameter("learning", "alpha"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if self.gamma, err = Float64Parameter("learning", "gamma"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if self.lambda, err = Float64Parameter("learning", "lambda"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if self.epsilon, err = Float64Parameter("learning", "epsilon"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Return the index of the best action from a given state
func (self *QLearning) ArgmaxAction(s State) (indexOfBest uint, valueOfBest float64) {
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
func (self *QLearning) RandomAction(s State) (indexOfBest uint, valueOfBest float64) {
	indexOfBest = uint(rand.Intn(len(self.Q[s.Id])))
	valueOfBest = self.Q[s.Id][indexOfBest]
	return
}

// Return an epsilon-greedy action, its estimated value, and whether it was chosen greedily
func (self *QLearning) EpsilonGreedyAction(s State, epsilon float64) (indexOfBest uint, valueOfBest float64, wasGreedy bool) {
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
func (self *QLearning) Learn(env Environment) {
	for epoch := uint(1); epoch <= self.maxEpochs; epoch++ {
		// fmt.Println("*************************************************")
		// fmt.Printf( "* starting epoch %v\n", epoch)
		// fmt.Println("*************************************************")
		env.Reset()
		s := env.StartState()
		self.DiscretizeState(&s)
		numSteps := 0
		for !env.AtGoalState(s) && !env.AtFailState(s) {
			// fmt.Printf("s:  %v\n", s)

			// select an action
			aIndex, _, aGreedy := self.EpsilonGreedyAction(s, self.epsilon)
			a := self.actions[aIndex]
			// fmt.Printf("a:  %v\n", a)

			// observe reward, next state
			sp, reward := env.ApplyAction(s, a)
			self.DiscretizeState(&sp)
			// fmt.Printf("r:  %v\n", reward)
			// fmt.Printf("s': %v\n", sp)

			// calculate optimum value of next action
			_, argmaxAP := self.ArgmaxAction(sp)
			
			// calculate error
			delta := reward + self.gamma * argmaxAP - self.Q[s.Id][a.Id]
			//fmt.Printf("delta: %v\n", delta)

			// fmt.Printf("updating Q[%v,%v] from %v to %v\n", s.Id, a.Id, self.Q[s.Id][a.Id],
			// 	self.Q[s.Id][a.Id] + self.alpha * delta)
			//self.Q[s.Id][a.Id] += self.alpha * delta

			// // set the currently visited state in the trace
			if aGreedy {
				self.E[s.Id][a.Id] = 1.0
			} else {
				self.E[s.Id][a.Id] = 0.0
			}

			// // update the policy
			for i := range self.Q {
				for j := range self.Q[i] {
					self.Q[i][j] += self.alpha * delta * self.E[i][j]
					self.E[i][j] *= self.gamma * self.lambda
				}
			}

			// print the Q-values
			// for i := range self.Q {
			// 	fmt.Printf("state %v (%v):\t\t%v\n", i, self.states[i], self.Q[i])
			// }
			// fmt.Println("")

			// iterate the policy
			s = sp

			numSteps++
		}
		self.epsilon *= 0.95
		fmt.Printf("Epoch: %v -- Pole balanced for %v steps.\n", epoch, numSteps)
	}
}

// given an arbitrary state vector, set its id to that of the nearest state in the space
func (self *QLearning) DiscretizeState(s *State) {
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

// func (self *QLearning) SavePolicy(filename string) error {
// 	f, err := os.Create(filename)
// 	if err != nil {
// 		("error saving policy: could not open '%v' for writing.\n", filename)
// 	}
// }