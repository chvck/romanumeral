package romanumeral

import (
	"fmt"
	"strconv"
)

// Runner is called from main and is resposible for running the application
type Runner struct {
	Generator IRomanNumeralGenerator
}

// NewRunner initializes and returns a new Runner
func NewRunner() *Runner {
	r := &Runner{}
	r.Generator = NewRomanNumeralGenerator()

	return r
}

// Run runs the application
func (r *Runner) Run(args ...string) (*string, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("incorrect number of arguments passed")
	}

	if number, err := strconv.Atoi(args[0]); err != nil {
		return nil, fmt.Errorf("number passed must be an integer number")
	} else {
		roman, err := r.Generator.Generate(number)
		return &roman, err
	}
}
