package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	configFile := flag.String("conf", "", "A configuration file defining parameters of the run.")
	//policyFile := flag.String("policy", "policy.dat", "File containing a saved policy that will be used to initialize the learner.")
	flag.Parse()

	if *configFile == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := InitConfig(*configFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	env := CreateEnvironment()
	lrn := CreateLearner()

	lrn.Init(env)
	lrn.Learn(env)
	lrn.FollowPolicy(env)
	// fmt.Println(lrn)
	// fmt.Println(env.Features())
}
