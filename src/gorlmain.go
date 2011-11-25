package main

import (
//	"fmt"
	"flag"
//	conf "goconf.googlecode.com/hg"
)

func main() {
	var configFile string
	var policyFile string

	flag.StringVar(&configFile, "conf", "gorl.cfg", "A configuration file defining parameters of the run.")
	flag.StringVar(&policyFile, "policy", "policy.dat", "File containing a saved policy that will be used to initialize the learner.")
	flag.Parse()

	// if conf.Lookup("conf") == nil {
	// 	fmt.Println("required parameter 'conf' not specified.")
	// 	os.Exit(1)
	// }
	
	InitConfig(configFile)

	// dims := make([][]float64, len(myDP))
	// for i := 0; i < len(myDP); i++ {
	// 	dims[i] = Linspace(-1.0, 1.0, int(myDP[i]))
	// }
	// states := BuildLattice(dims)

	// fmt.Println("")
	// for i := range states {
	// 	fmt.Println(states[i])
	// }
}
