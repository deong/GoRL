package environment

import (
	. "core"
	"os"
	"fmt"
)

// Each parameter has a min and max value.
type Range struct {
	Min float64
	Max float64
}

// Defines an interface for all reinforcement learning problems
type Environment interface {
	Features() []Range
	ApplyAction(s State, a Action) (sp State, reward float64)
	AtGoalState(s State) bool
	AtFailState(s State) bool
}

// return a new reinforcement learning environment
func CreateEnvironment() Environment {
	var name string
	var err os.Error
	if name, err = StringParameter("environment", "problem"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if name == "cart_pole" {
		return new(CartPoleEnv)
	} else if name == "mountain_car" {
		return new(MountainCarEnv)
	}
	return nil
}
