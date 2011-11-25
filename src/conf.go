// Defines a convenient interface for configuration information.
//
// Hides the details of parsing out the parameters and provides additional 
// functionality (e.g., parsing an array of int or float64 values.

package main

import (
	"fmt"
	gc "goconf.googlecode.com/hg"
	"strings"
	"strconv"
	"os"
)

// private variable that stores the parsed configuration info
var (
	rc *gc.ConfigFile
)

// parse the configuration file, returning any error encountered
func InitConfig(file string) (err os.Error) {
	rc, err = gc.ReadConfigFile(file)
	return 
}

// return the value of a given parameter as a string
func StringParameter(sec, name string) (val string, err os.Error) {
	val, err = rc.GetString(sec, name)
	return
}

// return the value of the parameter as an int
func IntParameter(sec, name string) (val int, err os.Error) {
	val, err = rc.GetInt(sec, name)
	return 
}

// return the value of the parameter as an int
func Float64Parameter(sec, name string) (val float64, err os.Error) {
	val, err = rc.GetFloat64(sec, name)
	return 
}

// return the value of the parameter as an int
func IntArrayParameter(sec, name string) (val []int, err os.Error) {
	var s string
	if s, err = rc.GetString(sec, name); err != nil {
		val = []int{}
	} else {
		val = parseIntVector(s, " ")
	}
	return 
}

// return the value of the parameter as an int
func Float64ArrayParameter(sec, name string) (val []float64, err os.Error) {
	var s string
	if s, err = rc.GetString(sec, name); err != nil {
		val = []float64{}
	} else {
		val = parseFloat64Vector(s, " ")
	}
	return 
}

// parse a string of numbers separated by spaces into a slice of ints
func parseFloat64Vector(str, sep string) []float64 {
	tokens := strings.Split(str, sep)
	vals := make([]float64, len(tokens))
	for i := range tokens {
		val, err := strconv.Atof64(tokens[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			vals[i] = val
		}
	}
	return vals
}

// parse a string of numbers separated by spaces into a slice of float64s 
func parseIntVector(str, sep string) []int {
	tokens := strings.Split(str, sep)
	vals := make([]int, len(tokens))
	for i := range tokens {
		val, err := strconv.Atoi(tokens[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			vals[i] = val
		}
	}
	return vals
}
