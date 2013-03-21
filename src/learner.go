package main

import (
	"os"
	"fmt"
)

type Learner interface {
	Init(env Environment)
	Learn(env Environment)
	ArgmaxAction(s State) (indexOfBest uint, valueOfBest float64)
	RandomAction(s State) (indexOfBest uint, valueOfBest float64)
	EpsilonGreedyAction(s State, epsilon float64) (indexOfBest uint, valueOfBest float64, wasGreedy bool)
	FollowPolicy(env Environment)
}

func CreateLearner () Learner {
	var name string
	var err error
	if name, err = StringParameter("learning", "learner"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if name == "qlearning" {
		return new(QLearning)
	} else if name == "rlearning" {
		return new(RLearning)
	}
	return nil
}