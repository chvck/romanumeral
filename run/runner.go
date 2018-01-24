package run

import (
	"fmt"
	"strconv"

	"github.com/chvck/romanumeral"
)

// Runner is called from main and is resposible for running the application
type Runner struct {
	generator romanumeral.IRomanNumeralGenerator
}

// NewRunnerWithGenerator initializes and returns a new Runner, using
// the specified generator
func NewRunnerWithGenerator(generator romanumeral.IRomanNumeralGenerator) *Runner {
	r := &Runner{}
	r.generator = generator

	return r
}

// NewRunner initializes and returns a new Runner
func NewRunner() *Runner {
	r := &Runner{}
	r.generator = romanumeral.NewRomanNumeralGenerator()

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
		roman, err := r.generator.Generate(number)
		return &roman, err
	}
}
