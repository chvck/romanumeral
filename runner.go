package romanumeral

import (
	"fmt"
)

// Runner is called from main and is resposible for running the application
type Runner struct {
}

// Run runs the application
func (r *Runner) Run(args ...interface{}) (*int, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("incorrect number of arguments passed")
	}

	if _, ok := args[0].(int); !ok {
		return nil, fmt.Errorf("number passed must be an integer number")
	} else {
		num := 0
		return &num, nil
	}
}
