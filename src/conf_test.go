package main

import (
	"testing"
)

func TestInitConfig(t *testing.T) {
	// check an existing file
	if err := InitConfig("sample_test.cfg"); err != nil {
		t.Errorf("Error opening configuration file: %v\n", "sample.cfg")
	} else {
		// check a few parameters
		if lm, err2 := Float64Parameter("learning", "lambda"); err2 != nil {
			t.Error("Error reading parameter lambda.\n")
		} else if !epsilonEqual(lm, 0.9, 0.00001) {
			t.Errorf("Incorrect value of parameter lambda (%v) found: expected 0.9.\n", lm)
		}

		// check an int parameter
		if eps, err3 := IntParameter("learning", "epochs"); err3 != nil {
			t.Error("Error reading parameter epochs.\n")
		} else if eps != 200 {
			t.Errorf("Incorrect value of parameter epochs (%v) found: expected 200.\n", eps)
		}
		
		// test the string retrieval
		if ln, err4 := StringParameter("learning", "learner"); err4 != nil {
			t.Error("Error reading parameter learner.\n")
		} else if ln != "qlearning" {
			t.Errorf("Incorrect value of parameter learner (%v) found: expected 'qlearning'.\n", ln)
		}

		// check a parameter that's misspelled
		if _, err4 := StringParameter("learning", "learnner"); err4 == nil {
			t.Error("Expected to get an error reading parameter learnner.\n")
		}

		// try an []int parameter
		if grid, err5 := IntArrayParameter("environment", "state_grid"); err5 != nil {
			t.Error("Error reading parameter state_grid.\n")
		} else {
			right := []int{5, 5, 5, 5}
			for i := range grid {
				if grid[i] != right[i] {
					t.Errorf("Incorrect value of parameter state_grid (%v) found: expected %v.\n",
						grid, right)
				}
			}
		}
	}
}
