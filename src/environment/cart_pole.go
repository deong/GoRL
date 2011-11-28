package environment

import (
	"math"
)

type CartPoleEnv struct {
	steps uint
}

var featureRanges = []Range{
	Range{-4.5, 4.5},
	Range{-3.0, 3.0},
	Range{-math.Pi / 4.0, math.Pi / 4.0},
	Range{-math.Pi / 4.0, math.Pi / 4.0},
}

func (env *CartPoleEnv) Features() (f []Range) {
	f = featureRanges
	return
}
